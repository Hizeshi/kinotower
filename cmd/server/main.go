package main

import (
	"log/slog"
	"os"

	core_server "github.com/Hizeshi/kinotower/internal/core/server"
)

func main() {
	slog.Info("Starting server...")

	server := core_server.NewServer()

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}