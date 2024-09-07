package handlers

import (
	"net/http"
	. "rwa/internal/controller"

	"time"

	auth "rwa/internal/handlers/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// nolint:errcheck
func GetApp() http.Handler {

	r := chi.NewRouter()
	// -------    middleware   -------
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(auth.AuthMiddleware)
	// ---------   routers     -------
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	r.Mount("/api", buildEndpoints())
	return r
}

func buildEndpoints() chi.Router {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		userController := NewUserController()
		r.Post("/login", userController.Login)
		r.Get("/", userController.GetCurrent)
		r.Put("/", userController.UpdateCurrent)
		r.Post("/", userController.Register)
	})

	r.Mount("/tags", func() http.Handler {
		r := chi.NewRouter()
		tagsController := NewTagsController()
		r.Get("/", tagsController.ListTags)
		return r
	}())

	r.Mount("/articles", func() http.Handler {
		r := chi.NewRouter()
		articleController := NewArticleService()
		r.Get("/", articleController.ServeHTTP)
		return r
	}())

	return r
}
