package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// nolint:errcheck
func GetApp() http.Handler {

	r := chi.NewRouter()
	// -------    middleware   -------
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(AuthMiddleware)
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
		userHandler := NewUserHandler()
		r.Post("/login", userHandler.Login)
		r.Get("/", userHandler.GetCurrent)
		r.Put("/", userHandler.UpdateCurrent)
		r.Post("/", userHandler.Register)
	})

	r.Mount("/tags", func() http.Handler {
		r := chi.NewRouter()
		tagsHandler := NewTagsHandler()
		r.Get("/", tagsHandler.ListTags)
		return r
	}())

	r.Mount("/articles", func() http.Handler {
		r := chi.NewRouter()
		articleHandler := NewArticleHandler()
		r.Get("/", articleHandler.GetRecent)
		r.Post("/", articleHandler.Create)
		r.Get("/feed", articleHandler.GetByFollowers)
		r.Put("/{slug}", articleHandler.Update)
		r.Get("/{slug}", articleHandler.GetArticle)
		r.Delete("/{slug}", articleHandler.DeleteArticle)
		r.Get("/{slug}/comments", articleHandler.GetComments)
		r.Post("/{slug}/comments", articleHandler.PostComments)
		r.Delete("/{slug}/comments/{id}", articleHandler.DeleteComments)

		return r
	}())

	return r
}
