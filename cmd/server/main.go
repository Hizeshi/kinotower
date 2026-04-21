package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	core_db "github.com/Hizeshi/kinotower/internal/core/db"

	core_server "github.com/Hizeshi/kinotower/internal/core/server"
)

func main() {
	slog.Info("Starting server...")

	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found, retying on environment variables")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		slog.Error("SUPABASE_URL or SUPABASE_KEY is not set")
		os.Exit(1)
	}

	supabaseClient := core_db.NewSupabaseClient(supabaseURL, supabaseKey)
	slog.Info("Supabase client initialized")

	server := core_server.NewServer(supabaseClient)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}