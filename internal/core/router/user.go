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

	ratingHandler "github.com/Hizeshi/kinotower/internal/features/ratings/handler"
	ratingRepo "github.com/Hizeshi/kinotower/internal/features/ratings/repository"
	ratingService "github.com/Hizeshi/kinotower/internal/features/ratings/servise"
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
		router.Use(middleware.RequareAuth) 
		
		router.Post("/", revHndl.Create)     
		router.Get("/", revHndl.GetUserReviews) 
		router.Delete("/{id}", revHndl.Delete) 
	})

	rateRepo := ratingRepo.NewRatingRepository(r.supabaseClient)
	rateSvc := ratingService.NewRatingService(rateRepo)
	rateHndl := ratingHandler.NewRatingHandler(rateSvc)

	router.Route("/{user-id}/ratings", func(router chi.Router) {
		router.Use(middleware.RequareAuth)
		
		router.Post("/", rateHndl.Create)
		router.Get("/", rateHndl.Get)
		router.Delete("/{id}", rateHndl.Delete)
	})

	return router
}