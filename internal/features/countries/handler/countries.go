package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Hizeshi/kinotower/internal/core/domain"
	"github.com/Hizeshi/kinotower/internal/features/countries/servise"
)

type CountryHandler struct {
	service *servise.CountryService
}

func NewCountryHandler(service *servise.CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (h *CountryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	countries, err := h.service.GetAllCountries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string][]domain.Country{
		"countries": countries,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}