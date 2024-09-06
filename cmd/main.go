package main

import (
	"fmt"
	"net/http"
	"rwa/internal/handlers"
)

func main() {
	addr := ":8080"
	h := handlers.GetApp() // main handler with new MUX roter since golang 1.22
	fmt.Println("start server at", addr)

	//nolint:errcheck
	http.ListenAndServe(addr, h)
}
