package handler

import (
	"context"
	"log"
	"net/http"
)

func ShutdownHandler(cancel context.CancelFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("shutdown request received")
		// cancel context
		cancel()
	})
}
