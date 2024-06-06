package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ropehapi/desafio-stone/internal/application/usecase"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"net/http"
)

type WebPersonHandler struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewWebPersonHandler(OrderRepository entity.PersonRepositoryInterface) *WebPersonHandler {
	return &WebPersonHandler{
		PersonRepository: OrderRepository,
	}
}

func (h *WebPersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreatePersonUseCaseInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createPersonUsecase := usecase.NewCreatePersonUseCase(h.PersonRepository)

	output, err := createPersonUsecase.Execute(dto)
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

	getPersonUsecase := usecase.NewGetPersonUseCase(h.PersonRepository)

	output, err := getPersonUsecase.Execute(dto)
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

func (h *WebPersonHandler) List(w http.ResponseWriter, r *http.Request) {
	listPersonUsecase := usecase.NewListPersonUseCase(h.PersonRepository)

	output, err := listPersonUsecase.Execute()
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

func (h *WebPersonHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreatePersonUseCaseInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatePersonUsecase := usecase.NewUpdatePersonUsecase(h.PersonRepository)
	output, err := updatePersonUsecase.Execute(chi.URLParam(r, "id"), dto)
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

func (h *WebPersonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	deletePersonUsecase := usecase.NewDeletePersonUseCase(h.PersonRepository)

	err := deletePersonUsecase.Execute(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
