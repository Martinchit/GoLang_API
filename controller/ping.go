package controller

import (
	"encoding/json"
	"net/http"

	"github.com/first-api/view"
)

// Ping Func
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getPing(w)
		}
	}
}

// GetPing Func
func getPing(w http.ResponseWriter) {
	res := view.PingAPIResponse{
		Code: http.StatusOK,
		Body: "Pong",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
