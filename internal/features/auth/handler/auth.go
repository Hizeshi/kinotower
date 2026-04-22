package handler

import (
	"encoding/json"
	"net/http"
	"log/slog"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/auth/servise"
)

type AuthHandler struct {
	service *servise.AuthServise
}

func NewAuthHandler(service *servise.AuthServise) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req domain.SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.SignUp(r.Context(), req)
	if err != nil {
		slog.Error("Error during signup", "error", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) 
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req domain.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.SignIn(r.Context(), req)
	if err != nil {
		slog.Warn("Failed login attempt", "email", req.Email)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "invalid",
			"message": "Wrong email or password",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}
