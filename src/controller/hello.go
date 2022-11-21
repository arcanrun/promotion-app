package controller

import (
	"encoding/json"
	"github.com/promotionApp/src/repository"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(repository.Promotions)
}
