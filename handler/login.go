package handler

import (
	"encoding/json"
	"github.com/eltaljohn/crudapi/authorization"
	"github.com/eltaljohn/crudapi/model"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "no valid structure", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}
	if !isLoginValid(&data) {
		resp := newResponse(Error, "no valid credentials", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "could not generate token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	response := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusCreated, response)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contact@ed.team" && data.Password == "123456"
}
