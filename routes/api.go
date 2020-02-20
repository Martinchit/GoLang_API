package routes

import (
	"net/http"

	"github.com/first-api/controller"
)

// Register the mux
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", controller.Ping())
	mux.HandleFunc("/api/entry", controller.Todo())
	mux.HandleFunc("/api/entry/", controller.Todo())
	return mux
}
