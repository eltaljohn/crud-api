package handler

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
	Data        any    `json:"data"`
}

func newResponse(messageType, message string, data any) response {
	return response{
		messageType,
		message,
		data,
	}
}

func responseJSON(w http.ResponseWriter, statusCode int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
