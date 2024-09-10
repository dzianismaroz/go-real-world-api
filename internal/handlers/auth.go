package handlers

import (
	"context"
	"net/http"
	repository "rwa/internal/repository/inmemory"
	"rwa/internal/utils"
)

// permitted paths :
// * POST /api/users        Content-Type: application/json
// * POST /api/users/login  Content-Type: application/json
// * GET  /api/tags
func isPermitted(req *http.Request) bool {
	method := req.Method
	path := req.URL.Path
	// log.Printf("resolving : method %s path %s", method, path)
	switch {
	case method == http.MethodPost && (path == "/api/users" || path == "/api/users/login"):
		fallthrough
	case method == http.MethodGet && (path == "/api/tags" || path == "/api/articles"):
		return true
	default:
		return false
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if isPermitted(req) {
			next.ServeHTTP(rw, req)
			return
		}
		// Authorize request. Fail if not allowed
		userId, err := repository.GetSessionManager().Check(req)
		if err != nil {
			http.Error(rw, "unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(req.Context(), utils.UserKey, userId)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})

}
