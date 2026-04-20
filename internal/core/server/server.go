package core_server

import (
	"net/http"

	core_router "github.com/Hizeshi/kinotower/internal/core/router"
)

type Server struct {
	http.Server

}

func NewServer() *Server {
    router := core_router.NewRouter()
	
    return &Server{
		Server: http.Server{
			Addr:":3000",
			Handler: router.RegisterRoutes(),
		},
	}

}

