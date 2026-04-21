package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Hizeshi/kinotower/internal/features/genders/servise"
)

type GenderHandler struct {
	service *servise.GenderService
}

func NewGenderHandler(service *servise.GenderService) *GenderHandler {
	return &GenderHandler{service: service}
}

func (h *GenderHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	genders, err := h.service.GetAllGenders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"genders": genders,
	}

	w.Header().Set("Content-Type", "application/json")
	
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}