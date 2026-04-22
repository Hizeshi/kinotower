package core_router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nedpals/supabase-go"
)

type Router struct {
	supabaseClient *supabase.Client
}

func NewRouter(supabaseClient *supabase.Client) *Router {
	return &Router{
		supabaseClient: supabaseClient,
	}
}

func (r *Router) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Route("/api/v1", func(rl chi.Router) {
		rl.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
		rl.Mount("/films", r.filmRoutes())
		rl.Mount("/auth", r.authRoutes())
		rl.Mount("/genders", r.genderRoutes())
		rl.Mount("/countries", r.countryRoutes())
		rl.Mount("/categories", r.categoryRoutes())
		rl.Mount("/users", r.userRoutes())
	})
	return router
}
