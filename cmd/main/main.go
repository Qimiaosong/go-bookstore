package main

import (
	"log"
	"net/http"
	//"github.com/akhil/go-bookstore/pkg/routes"
	"github.com/Qimiaosong/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
