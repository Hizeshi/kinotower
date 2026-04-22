package core_router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Hizeshi/kinotower/internal/features/films/handler"
	"github.com/Hizeshi/kinotower/internal/features/films/repository"
	"github.com/Hizeshi/kinotower/internal/features/films/servise"
)

func (r *Router) filmRoutes() http.Handler {
	router := chi.NewRouter()

	repo := repository.NewFilmRepository(r.supabaseClient)
	svc := servise.NewFilmService(repo)
	hndl := handler.NewFilmHandler(svc)

	router.Get("/", hndl.GetAll)
	router.Get("/{id}", hndl.GetByID)
	router.Get("/{id}/reviews", hndl.GetReviews)

	return router
}