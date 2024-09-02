package main

import (
	"fmt"
	"net/http"
	"rwa/internal/handlers"
)

func main() {
	addr := ":8080"
	h := handlers.GetApp()
	fmt.Println("start server at", addr)
	//nolint:errcheck
	http.ListenAndServe(addr, h)
}
