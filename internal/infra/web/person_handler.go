package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ropehapi/desafio-stone/internal/application/usecase"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"net/http"
)

type WebPersonHandler struct {
	PersonRepositoty entity.PersonRepositoryInterface
}

func NewWebPersonHandler(OrderRepository entity.PersonRepositoryInterface) *WebPersonHandler {
	return &WebPersonHandler{
		PersonRepositoty: OrderRepository,
	}
}

func (h *WebPersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreatePersonUseCaseInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createPersonUsecase := usecase.NewCreatePersonUseCase(h.PersonRepositoty)

	output, err := createPersonUsecase.Exec(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebPersonHandler) Get(w http.ResponseWriter, r *http.Request) {
	dto := usecase.GetPersonUseCaseInputDTO{
		ID: chi.URLParam(r, "id"),
	}

	getPersonUsecase := usecase.NewGetPersonUseCase(h.PersonRepositoty)

	output, err := getPersonUsecase.Exec(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
