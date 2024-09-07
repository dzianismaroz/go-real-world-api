package main

import (
	"log"
	"net/http"
	"rwa/internal/handlers"
)

func main() {
	addr := ":8080"
	h := handlers.GetApp() // main handler with go-Chi router
	log.Println("start server at", addr)

	//nolint:errcheck
	http.ListenAndServe(addr, h)
}
