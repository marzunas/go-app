package handler

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"github.com/marzunas/go-app/conf"
	"github.com/marzunas/go-app/store"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var idCounter = new(int64)
var cfg = conf.GetConfig()

func HashHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// reject requests after shutdown initiation
		if checkShutdown(ctx, w) {
			log.Println("server shutting down")
			return
		}
		password := r.FormValue("password")
		if password != "" {
			log.Println("hash compute request")
			var id = getNextId()
			computeHash(id, password)
			// return HTTP 202 request ACK
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(strconv.Itoa(id)))
			w.Write([]byte("\n"))
			log.Println("hash compute request processed")
		} else {
			log.Println("hash fetch request")
			var reqId = r.URL.Path[len("/hash/"):]
			if reqId != "" {
				id, err := strconv.Atoi(reqId)
				if err != nil {
					// un parsable request id
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("un-parsable request id: " + reqId))
					w.Write([]byte("\n"))
					return
				}
				// return HTTP 200 with hash
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(store.HashStore[id]))
				w.Write([]byte("\n"))
			}
			log.Println("hash fetch request processed")
		}
	})
}

func checkShutdown(ctx context.Context, w http.ResponseWriter) bool {
	if ctx.Err() != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("unable to process request - server shutting down"))
		return true
	}
	return false
}

func getNextId() int {
	// atomically increment counter
	atomic.AddInt64(idCounter, 1)
	return int(*idCounter)
}

func computeHash(id int, password string) {
	// time delayed function to compute hash
	time.AfterFunc(time.Duration(cfg.HashWaitPeriod)*time.Second, func() {
		log.Printf("computing hash for req: %v ", id)
		var hash = sha512.New()
		hash.Write([]byte(password))
		store.HashStore[id] = base64.StdEncoding.EncodeToString(hash.Sum(nil))
	})
}
