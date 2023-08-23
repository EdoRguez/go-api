package main

import (
	"net/http"

	"github.com/edorguez/go-api/db"
	"github.com/edorguez/go-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
