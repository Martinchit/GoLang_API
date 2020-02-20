package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/first-api/model"
	"github.com/first-api/view"
)

// Entry point
func Todo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTodo(w, r)
		case http.MethodGet:
			name := r.URL.Query().Get("name")
			if name == "" {
				getTodos(w)
			} else {
				getTodo(w, name)
			}
		case http.MethodDelete:
			deleteTodo(w, r)
		}
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	data := view.Todo{}
	json.NewDecoder(r.Body).Decode(&data)
	name, todo := data.Name, data.Todo
	err := model.CreateTodo(name, todo)
	if err != nil {
		w.Write([]byte("Insert Error"))
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func getTodos(w http.ResponseWriter) {
	data, err := model.GetTodos()
	if err != nil {
		w.Write([]byte("Read Error"))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func getTodo(w http.ResponseWriter, name string) {
	data, err := model.GetTodo(name)
	if err != nil {
		w.Write([]byte("Read Error"))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	name := paths[len(paths)-1]
	err := model.DeleteTodo(name)
	if err != nil {
		w.Write([]byte("Insert Error"))
	}
	res := view.TodoAPIResponse{
		Status: "Item Deleted",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
