package server

import (
	"context"
	"fmt"
	"github.com/marzunas/go-app/conf"
	"github.com/marzunas/go-app/handler"
	"github.com/marzunas/go-app/store"
	"log"
	"net/http"
	"strconv"
	"time"
)

var cfg = conf.GetConfig()

// interceptor to log request and response
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("request: " + r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

// interceptor to track execution times
func execTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		elapsed := time.Now().Sub(start)
		store.ExecTimes = append(store.ExecTimes, elapsed.Microseconds())
	})
}

func Boot() {
	log.Println("initializing http server")
	mux := http.NewServeMux()
	server := &http.Server{Addr: ":" + strconv.Itoa(cfg.Port), Handler: mux}
	// create cancel hook
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux.Handle("/hash", logger(execTimer(handler.HashHandler(ctx))))
	// optimize this to bind handler with wild card /hash/* pattern
	mux.Handle("/hash/", logger(handler.HashHandler(ctx)))
	mux.Handle("/stats", logger(handler.StatsHandler()))
	mux.Handle("/shutdown", handler.ShutdownHandler(cancel))
	fmt.Printf("starting server at port 8080\n")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	select {
	case <-ctx.Done():
		log.Printf("waiting for %v routines to complete", 2)
		time.Sleep(time.Duration(cfg.HashWaitPeriod) * time.Second)
		server.Shutdown(ctx)
	}
	log.Println("server stopped")
}
