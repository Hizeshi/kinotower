package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/reviews/servise"
	"github.com/go-chi/chi/v5"
)

type ReviewHandler struct {
	service *servise.ReviewService
}

func NewReviewHandler(service *servise.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "user-id"))

	var req domain.CreateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	review, err := h.service.AddReview(r.Context(), userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(review)
}

func (h *ReviewHandler) GetUserReviews(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "user-id"))

	reviews, err := h.service.GetUserReviews(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"reviews": reviews, 
	})
}

func (h *ReviewHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "user-id"))
	reviewID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteReview(r.Context(), userID, reviewID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		message := "Review not found"
		if err.Error() == "User not found" {
			message = "User not found"
		}
		json.NewEncoder(w).Encode(map[string]string{"message": message})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}