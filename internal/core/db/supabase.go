package db

import "github.com/nedpals/supabase-go"

func NewSupabaseClient(supabaseURL, supabaseKey string) *supabase.Client {
	client := supabase.CreateClient(supabaseURL, supabaseKey)
	return client
}