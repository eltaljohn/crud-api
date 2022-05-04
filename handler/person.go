package handler

import (
	"encoding/json"
	"errors"
	"github.com/eltaljohn/crudapi/model"
	"net/http"
	"strconv"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "body malformed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "error creating person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "person created successfully", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadGateway, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "error getting all people", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "the id must be int", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "body malformed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "error updating person", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "person updated successfully", nil)
	responseJSON(w, http.StatusBadRequest, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "the id must be int", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "user id does not exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "error deleting person", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}
