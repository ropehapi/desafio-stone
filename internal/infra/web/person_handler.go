package web

import (
	"encoding/json"
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

	createOrder := usecase.NewCreatePersonUseCase(h.PersonRepositoty)
	output, err := createOrder.Exec(dto)
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
