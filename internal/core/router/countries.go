package core_router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	repository "github.com/Hizeshi/kinotower/internal/features/countries/repository"
	servise "github.com/Hizeshi/kinotower/internal/features/countries/servise"
	handler "github.com/Hizeshi/kinotower/internal/features/countries/handler"
)

func (r *Router) countryRoutes() http.Handler {
	router := chi.NewRouter()

	repo := repository.NewCountryRepository(r.supabaseClient)
	svc := servise.NewCountryService(repo)
	hndl := handler.NewCountryHandler(svc)

	router.Get("/", hndl.GetAll)

	return router
}
