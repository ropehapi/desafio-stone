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

// Create Person godoc
// @Summary Create person
// @Description Create person
// @Tags persons
// @Accept json
// @Produce json
// @Param request body usecase.CreateUpdatePersonUsecaseInputDTO true "person request"
// @Success 201
// @Failure 400
// @Failure 500 {object} Error
// @Router /person [post]
func (h *WebPersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateUpdatePersonUsecaseInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	createPersonUsecase := usecase.NewCreatePersonUseCase(h.PersonRepository)

	output, err := createPersonUsecase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}

// Get Person godoc
// @Summary Get person
// @Description Get person
// @Tags persons
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Success 200
// @Failure 500 {object} Error
// @Router /person/{id} [get]
func (h *WebPersonHandler) Get(w http.ResponseWriter, r *http.Request) {
	getPersonUsecase := usecase.NewGetPersonUseCase(h.PersonRepository)

	output, err := getPersonUsecase.Execute(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}

// List Person godoc
// @Summary List person
// @Description List person
// @Tags persons
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} Error
// @Router /person [get]
func (h *WebPersonHandler) List(w http.ResponseWriter, r *http.Request) {
	listPersonUsecase := usecase.NewListPersonUseCase(h.PersonRepository)

	output, err := listPersonUsecase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}

// Update Person godoc
// @Summary Update person
// @Description Update person
// @Tags persons
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Param request body usecase.CreateUpdatePersonUsecaseInputDTO true "person request"
// @Success 200
// @Failure 400
// @Failure 500 {object} Error
// @Router /person/{id} [put]
func (h *WebPersonHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateUpdatePersonUsecaseInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	updatePersonUsecase := usecase.NewUpdatePersonUsecase(h.PersonRepository)
	output, err := updatePersonUsecase.Execute(chi.URLParam(r, "id"), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}

// Delete Person godoc
// @Summary Delete person
// @Description Delete person
// @Tags persons
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Success 200
// @Failure 500 {object} Error
// @Router /person/{id} [delete]
func (h *WebPersonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	deletePersonUsecase := usecase.NewDeletePersonUseCase(h.PersonRepository)

	err := deletePersonUsecase.Execute(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}
