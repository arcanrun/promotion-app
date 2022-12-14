package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/promotionApp/src/dto"
	"github.com/promotionApp/src/service"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetAllPromotions(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(service.GetAllPromotions())
}

func AddPromotion(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	defer request.Body.Close()

	if body, err := io.ReadAll(request.Body); err != nil {
		log.Fatalln(err)
	} else {
		var incomingDto dto.PromotionDto

		json.Unmarshal(body, &incomingDto)

		err := service.AddPromotion(incomingDto)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func GetPromotionById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	defer request.Body.Close()

	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	if responseDto, err := service.GetPromotingById(int64(id)); err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(dto.ErrorDto{
			Message: err.Error(),
		})
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(responseDto)
	}
}

func UploadCSV(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	if file, _, err := request.FormFile("file"); err != nil {
		json.NewEncoder(writer).Encode(dto.ErrorDto{
			Message: err.Error(),
		})

		writer.WriteHeader(http.StatusBadRequest)
	} else {
		service.UploadCSV(file)

		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(dto.SimpleResponse{
			Message: "File uploading has been started",
		})
	}
}

func DeleteAllPromotions(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
	service.DeleteAllPromotions()
}
