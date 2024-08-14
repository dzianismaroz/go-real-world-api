package main

import (
	"net/http"
)

// сюда писать код

func GetApp() http.Handler {
	return &AppHandler{}
}

type AppHandler struct {
}

func (a *AppHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

}
