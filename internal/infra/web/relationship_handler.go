package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
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

func (h *WebRelationshipHandler) GetTree(w http.ResponseWriter, r *http.Request) {
	getTreeUsecase := usecase.NewGetPersonTreeUseCase(h.PersonRepository, h.RelationshipRepository)

	output, err := getTreeUsecase.Execute(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *WebRelationshipHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateRelationshipInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createRelationshipUsecase := usecase.NewCreateRelationshipUsecase(h.PersonRepository, h.RelationshipRepository)

	outputDto, err := createRelationshipUsecase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(outputDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
