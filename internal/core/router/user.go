package core_router

import (
	"net/http"

	"github.com/Hizeshi/kinotower/internal/core/middleware"
	userHandler "github.com/Hizeshi/kinotower/internal/features/users/handler"
	userRepo "github.com/Hizeshi/kinotower/internal/features/users/repository"
	userService "github.com/Hizeshi/kinotower/internal/features/users/servise"
	
	reviewHandler "github.com/Hizeshi/kinotower/internal/features/reviews/handler"
	reviewRepo "github.com/Hizeshi/kinotower/internal/features/reviews/repository"
	reviewService "github.com/Hizeshi/kinotower/internal/features/reviews/servise"
	"github.com/go-chi/chi/v5"
)

func (r *Router) userRoutes() http.Handler {
	router := chi.NewRouter()

	repo := userRepo.NewUserRepository(r.supabaseClient)
	svc := userService.NewUserService(repo)
	hndl := userHandler.NewUserHandler(svc)

	router.Use(middleware.RequareAuth)

	router.Get("/{id}", hndl.GetProfile)
	router.Put("/{id}", hndl.Update)
	router.Delete("/", hndl.Delete)

	revRepo := reviewRepo.NewReviewRepository(r.supabaseClient)
	revSvc := reviewService.NewReviewService(revRepo)
	revHndl := reviewHandler.NewReviewHandler(revSvc)

	router.Route("/{user-id}/reviews", func(router chi.Router) {
		router.Use(middleware.RequareAuth) // 
		
		router.Post("/", revHndl.Create)     // POST /api/v1/users/{user-id}/reviews
		router.Get("/", revHndl.GetUserReviews) // GET /api/v1/users/{user-id}/reviews 
		router.Delete("/{id}", revHndl.Delete)  // DELETE /api/v1/users/{user-id}/reviews/{id} 
	})

	return router
}