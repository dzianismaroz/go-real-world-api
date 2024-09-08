package handlers

import "net/http"

type TagHandler struct {
}

func NewTagsHandler() *TagHandler {
	return &TagHandler{}
}

func (h *TagHandler) ListTags(rw http.ResponseWriter, req *http.Request) {

}
