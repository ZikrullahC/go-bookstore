package main

import (
	"log"
	"net/http"

	"github.com/ZikrullahC/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3306", r))

}