package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ropehapi/desafio-stone/internal/application/usecase"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type WebRelationshipHandler struct {
	PersonRepository       entity.PersonRepositoryInterface
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewWebRelationshipHandler(personRepository entity.PersonRepositoryInterface, relationshipRepository entity.RelationshipRepositoryInterface) *WebRelationshipHandler {
	return &WebRelationshipHandler{
		PersonRepository:       personRepository,
		RelationshipRepository: relationshipRepository,
	}
}

// GetRelationshipsDescendant Relationship godoc
// @Summary GetRelationshipsDescendant relationship
// @Description GetRelationshipsDescendant relationship
// @Tags relationship
// @Accept json
// @Produce json
// @Param id path string true "relationship ID" Format(uuid)
// @Success 200
// @Failure 400
// @Failure 500 {object} Error
// @Router /relationship/{id}/asc [get]
func (h *WebRelationshipHandler) GetRelationshipsAscendant(w http.ResponseWriter, r *http.Request) {
	getPersonRelationshipsAscendantUsecase := usecase.NewGetPersonRelationshipsAscendantUsecase(h.PersonRepository, h.RelationshipRepository)

	output, err := getPersonRelationshipsAscendantUsecase.Execute(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}

// GetRelationshipsDescendant Relationship godoc
// @Summary GetRelationshipsDescendant relationship
// @Description GetRelationshipsDescendant relationship
// @Tags relationship
// @Accept json
// @Produce json
// @Param id path string true "relationship ID" Format(uuid)
// @Success 200
// @Failure 400
// @Failure 500 {object} Error
// @Router /relationship/{id}/desc [get]
func (h *WebRelationshipHandler) GetRelationshipsDescendant(w http.ResponseWriter, r *http.Request) {
	getPersonRelationshipsDescendantUsecase := usecase.NewGetPersonRelationshipsDescendantUsecase(h.PersonRepository, h.RelationshipRepository)

	output, err := getPersonRelationshipsDescendantUsecase.Execute(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}

// Create Relationsip godoc
// @Summary Create relationship
// @Description Create relationship
// @Tags relationship
// @Accept json
// @Produce json
// @Param request body usecase.CreateRelationshipInputDTO true "relationship request"
// @Success 201
// @Failure 400
// @Failure 500 {object} Error
// @Router /relationship [post]
func (h *WebRelationshipHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateRelationshipInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	createRelationshipUsecase := usecase.NewCreateRelationshipUsecase(h.PersonRepository, h.RelationshipRepository)

	err = createRelationshipUsecase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete Relationship godoc
// @Summary Delete relationship
// @Description Delete relationship
// @Tags relationship
// @Accept json
// @Produce json
// @Param request body usecase.DeleteRelationshipUsecaseInputDTO true "relationship request"
// @Success 200
// @Failure 400
// @Failure 500 {object} Error
// @Router /relationship [delete]
func (h *WebRelationshipHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var dto usecase.DeleteRelationshipUsecaseInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}

	deleteRelationshipUsecase := usecase.NewDeleteRelationshipUsecase(h.RelationshipRepository)

	err = deleteRelationshipUsecase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
}
