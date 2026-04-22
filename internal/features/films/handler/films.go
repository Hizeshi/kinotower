package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/films/servise"
	"github.com/go-chi/chi/v5"
)

type FilmHandler struct {
	service *servise.FilmService
}

func NewFilmHandler(service *servise.FilmService) *FilmHandler {
	return &FilmHandler{service: service}
}

func (h *FilmHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	parseInt := func(s string, def int) int {
		if s == "" {
			return def
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return def
		}
		return val
	}

	sortBy := query.Get("sortBy")
	if sortBy == "" {
		sortBy = "name"
	}

	sortDir := query.Get("sortDir")
	if sortDir != "desc" {
		sortDir = "asc"
	}

	filter := domain.FilmFilter{
		Page:     parseInt(query.Get("page"), 1),
		Size:     parseInt(query.Get("size"), 10),
		Category: parseInt(query.Get("category"), 0),
		Country:  parseInt(query.Get("country"), 0),
		SortBy:   sortBy,
		SortDir:  sortDir,
		Search:   query.Get("search"), 
	}

	slog.Info("Received GET /films request", 
		"page", filter.Page, 
		"size", filter.Size, 
		"sortBy", filter.SortBy,
		"sortDir", filter.SortDir,
	)

	response, err := h.service.GetFilms(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Successfully fetched films", 
		"response_page", response.Page,
		"response_size", response.Size,
		"response_total", response.Total,
		"films_count_in_array", len(response.Films),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Error encoding JSON response", "error", err)
	}
}

func (h *FilmHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid film ID format", http.StatusBadRequest)
		return
	}

	film, err := h.service.GetFilmByID(r.Context(), id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Film not found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(film)
}


func (h *FilmHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	filmID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid film ID", http.StatusBadRequest)
		return
	}

	reviews, err := h.service.GetFilmReviews(r.Context(), filmID)
	if err != nil {
		slog.Error("Failed to get film reviews", "error", err, "filmID", filmID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Film not found",
		})
		return
	}

	response := map[string]interface{}{
		"reviews": reviews,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}