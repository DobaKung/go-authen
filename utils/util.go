package utils

import (
	"encoding/json"
	"net/http"
)

func NewMessage(success bool, code int, message string, data interface{}) MessagePayload {
	return MessagePayload{Success: success, Code: code, Message: message, Data: data}
}

func Respond(w http.ResponseWriter, status int, data MessagePayload) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status) // http status code
	json.NewEncoder(w).Encode(data)
}
