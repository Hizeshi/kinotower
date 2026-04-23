package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/core/middleware"
	"github.com/Hizeshi/kinotower/internal/features/ratings/servise"
	"github.com/go-chi/chi/v5"
)

type RatingHandler struct {
	service *servise.RatingService
}

func NewRatingHandler(service *servise.RatingService) *RatingHandler {
	return &RatingHandler{service: service}
}

func (h *RatingHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "user-id"))
	
	tokenUserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok || tokenUserID != userID {
		http.Error(w, "Forbidden: you can only access your own data", http.StatusForbidden)
		return
	}
	
	var req domain.CreateRatingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Ball < 1 || req.Ball > 10 {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"message": "Ball must be between 1 and 10"})
        return
    }

	id, err := h.service.AddRating(r.Context(), userID, req)
	if err != nil {
		if err == servise.ErrScoreExist {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest) 
			json.NewEncoder(w).Encode(map[string]string{"message": "Score exist"})
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "status": "success"})
}

func (h *RatingHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "user-id"))

	tokenUserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok || tokenUserID != userID {
		http.Error(w, "Forbidden: you can only access your own data", http.StatusForbidden)
		return
	}

	ratings, err := h.service.GetUserRatings(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"ratings": ratings})
}

func (h *RatingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	userID, _ := strconv.Atoi(chi.URLParam(r, "user-id"))

	tokenUserID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok || tokenUserID != userID {
		http.Error(w, "Forbidden: you can only access your own data", http.StatusForbidden)
		return
	}

	h.service.DeleteRating(r.Context(), id)
	w.WriteHeader(http.StatusNoContent)
}
