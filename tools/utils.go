package tools

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func Respond(w http.ResponseWriter, r Response) {
	json.NewEncoder(w).Encode(r)
}
