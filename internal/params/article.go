package params

import (
	"net/http"
	"strconv"
)

const (
	tagParam        = "tag"
	authorParam     = "author"
	favouritedParma = "favourited"
	limitParam      = "limit"
	offsetParam     = "offset"
)

type FilterParams struct {
	Tag       string
	Author    string // username
	Favorited string
	Limit     uint // default is 20
	Offest    uint // default is 0
}

func (f FilterParams) HasAuthorFilter() bool {
	return len(f.Author) > 1
}

func (f FilterParams) HasTagFilter() bool {
	return len(f.Tag) > 1
}

func GetFilterParams(req *http.Request) FilterParams {
	query := req.URL.Query()
	var (
		tag, author, favorited string
		limit                  uint = 20
		offset                 uint = 0
	)
	switch {
	case query.Get(tagParam) != "":
		tag = query.Get(tagParam)
	case query.Get(authorParam) != "":
		author = query.Get(authorParam)
	case query.Get(favouritedParma) != "":
		favorited = query.Get(favouritedParma)
	case query.Get(limitParam) != "":
		numval, err := strconv.Atoi(query.Get(limitParam))
		if err == nil && numval > 0 {
			limit = uint(numval)
		}
	case query.Get(offsetParam) != "":
		numval, err := strconv.Atoi(query.Get(offsetParam))
		if err == nil && numval > 0 {
			limit = uint(numval)
		}
	}
	return FilterParams{Tag: tag, Author: author, Favorited: favorited, Limit: limit, Offest: offset}
}
