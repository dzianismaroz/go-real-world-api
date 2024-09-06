package handlers

import (
	"net/http"
	service "rwa/internal/controller"

	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// nolint:errcheck
func GetApp() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	r.Mount("/api", buildEndpoints())
	return r
}

func buildEndpoints() chi.Router {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		userController := service.NewUserController()
		r.Get("/", userController.GetCurrent)
		r.Put("/", userController.UpdateCurrent)
		r.Post("/", userController.Register)
		r.Post("/login", userController.Login)
	})

	r.Mount("/tags", func() http.Handler {
		r := chi.NewRouter()
		tagsController := service.NewTagsController()
		r.Get("/", tagsController.ListTags)
		return r
	}())

	r.Mount("/articles", func() http.Handler {
		r := chi.NewRouter()
		articleController := service.NewArticleService()
		r.Get("/", articleController.ServeHTTP)
		return r
	}())

	return r
}
