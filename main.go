package main

import (
	"github.com/gorilla/mux"
	"github.com/promotionApp/src/controller"
	"log"
	"net/http"
	"time"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/hello", controller.HelloHandler)

	log.Printf("Server started at %s", time.Now())

	http.ListenAndServe(":8080", router)
}
