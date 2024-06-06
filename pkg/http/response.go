package http

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
