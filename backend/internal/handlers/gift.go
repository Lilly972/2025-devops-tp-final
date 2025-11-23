package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anthonyquere/christmas-gift-list/internal/models"
	"github.com/anthonyquere/christmas-gift-list/internal/repository"
	"github.com/go-chi/chi/v5"
)

type GiftHandler struct {
	repo *repository.GiftRepository
}

func NewGiftHandler() *GiftHandler {
	return &GiftHandler{
		repo: repository.NewGiftRepository(),
	}
}

func (h *GiftHandler) GetByPersonID(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personId")

	gifts, err := h.repo.GetByPersonID(r.Context(), personID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gifts)
}

func (h *GiftHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.CreateGiftInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gift, err := h.repo.Create(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(gift)
}

func (h *GiftHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var input models.UpdateGiftInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gift, err := h.repo.Update(r.Context(), id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gift)
}

func (h *GiftHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.repo.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *GiftHandler) SelectGift(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personId")
	giftID := chi.URLParam(r, "giftId")

	gift, err := h.repo.SelectGift(r.Context(), personID, giftID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gift)
}
