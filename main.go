package main

import (
	"log"
	"net/http"

	"github.com/first-api/model"
	"github.com/first-api/routes"
)

func main() {
	mux := routes.Register()
	db := model.Connect()
	defer db.Close()
	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
