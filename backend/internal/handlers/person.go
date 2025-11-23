package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anthonyquere/christmas-gift-list/internal/models"
	"github.com/anthonyquere/christmas-gift-list/internal/repository"
	"github.com/go-chi/chi/v5"
)

type PersonHandler struct {
	repo *repository.PersonRepository
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler{
		repo: repository.NewPersonRepository(),
	}
}

func (h *PersonHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	people, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func (h *PersonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	person, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.CreatePersonInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	person, err := h.repo.Create(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.repo.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
