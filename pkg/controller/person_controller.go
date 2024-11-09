// pkg/controller/person_controller.go
package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amartinezh/sdk-db/pkg/model"
	"github.com/amartinezh/sdk-db/pkg/service"
	"github.com/gorilla/mux"
)

type PersonController struct {
	PersonService *service.PersonService
}

func (pc *PersonController) AddPerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := pc.PersonService.CreatePerson(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (pc *PersonController) ListPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := pc.PersonService.GetAllPersons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(persons); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := pc.PersonService.UpdatePerson(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (pc *PersonController) RemovePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := pc.PersonService.DeletePerson(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
