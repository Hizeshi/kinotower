package core_router

import (
	"net/http"

	"github.com/Hizeshi/kinotower/internal/core/middleware"
	"github.com/Hizeshi/kinotower/internal/features/auth/handler"
	"github.com/Hizeshi/kinotower/internal/features/auth/repository"
	"github.com/Hizeshi/kinotower/internal/features/auth/servise"
	"github.com/go-chi/chi/v5"
)

func (r *Router) authRoutes() http.Handler {
	router := chi.NewRouter()

	repo := repository.NewAuthRepository(r.supabaseClient)
	svc := servise.NewAuthServise(repo)
	hndl := handler.NewAuthHandler(svc)

	router.Post("/signup", hndl.SignUp)
	router.Post("/signin", hndl.SignIn)

	router.Group(func(rl chi.Router) {
		rl.Use(middleware.RequareAuth)
		rl.Post("/signout", hndl.SignOut)
	})

	return router
}