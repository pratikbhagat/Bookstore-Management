package main

import (
	"log"
	"net/http"

	"Bookstore-Management/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9080", r))
}
