package core_router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Hizeshi/kinotower/internal/features/genders/handler"
	"github.com/Hizeshi/kinotower/internal/features/genders/repository"
	"github.com/Hizeshi/kinotower/internal/features/genders/servise"
)

func (r *Router) genderRoutes() chi.Router {
	routes := chi.NewRouter()

	repo := repository.NewGenderRepository(r.supabaseClient)
	svc := servise.NewGenderService(repo)
	hndl := handler.NewGenderHandler(svc)

	routes.Get("/", hndl.GetAll)

	return routes
}