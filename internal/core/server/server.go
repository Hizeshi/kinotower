package core_server

import (
	"net/http"

	core_router "github.com/Hizeshi/kinotower/internal/core/router"
	"github.com/nedpals/supabase-go"
)

type Server struct {
	http.Server

}

func NewServer(supabaseClient *supabase.Client) *Server {
    router := core_router.NewRouter(supabaseClient)
	
    return &Server{
		Server: http.Server{
			Addr:":3000",
			Handler: router.RegisterRoutes(),
		},
	}

}

