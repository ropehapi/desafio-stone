package web

import (
	"encoding/json"
	"github.com/ropehapi/desafio-stone/internal/application/usecase"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"net/http"
)

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

func (h *WebRelationshipHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateRelationshipInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createRelationshipUsecase := usecase.NewCreateRelationshipUsecase(h.PersonRepository, h.RelationshipRepository)

	outputDto, err := createRelationshipUsecase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(outputDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
