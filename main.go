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

	router.HandleFunc("/promotions", controller.GetAllPromotions).Methods(http.MethodGet)
	router.HandleFunc("/promotions", controller.AddPromotion).Methods(http.MethodPost)
	router.HandleFunc("/promotions/{id}", controller.GetPromotionById).Methods(http.MethodGet)

	log.Printf("Server started at %s", time.Now())

	http.ListenAndServe(":8080", router)
}
