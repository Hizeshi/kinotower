package handler

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/films/servise"

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

	filter := domain.FilmFilter{
		Page:     parseInt(query.Get("page"), 1),
		Size:     parseInt(query.Get("size"), 10),
		Category: parseInt(query.Get("category"), 0),
		Country:  parseInt(query.Get("country"), 0),
		SortBy:   query.Get("sortBy"),
		SortDir:  query.Get("sortDir"),
		Search:   query.Get("search"), 
	}

	response, err := h.service.GetFilms(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}