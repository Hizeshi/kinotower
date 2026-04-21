package core_router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	repository "github.com/Hizeshi/kinotower/internal/features/catregories/repository"
	servise "github.com/Hizeshi/kinotower/internal/features/catregories/servise"
	handler "github.com/Hizeshi/kinotower/internal/features/catregories/handler"
)

func (r *Router) categoryRoutes() http.Handler {
	router := chi.NewRouter()

	repo := repository.NewCategoryRepository(r.supabaseClient)
	svc := servise.NewCategoryService(repo)
	hndl := handler.NewCategoryHandler(svc)

	router.Get("/", hndl.GetAll)
	
	return router
}