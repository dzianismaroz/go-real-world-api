package handlers

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
	rw.WriteHeader(http.StatusNotFound)
	_, e := rw.Write([]byte("nothing!"))
	if e != nil {
		panic("somthign really wrong !")
	}

}
