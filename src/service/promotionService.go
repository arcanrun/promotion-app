package service

import (
	"encoding/csv"
	"github.com/google/uuid"
	"github.com/promotionApp/src/dto"
	"github.com/promotionApp/src/model"
	"github.com/promotionApp/src/repository"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"
)

func GetAllPromotions() []model.Promotion {
	return repository.GetAllPromotions()
}

func AddPromotion(dto dto.PromotionDto) error {
	promotion := model.Promotion{
		Id:             rand.Int63(),
		OriginalId:     dto.OriginalId,
		Price:          dto.Price,
		ExpirationDate: dto.ExpirationDate,
	}

	return repository.Save(promotion)
}

func GetPromotingById(id int64) (dto.PromotionResponseDto, error) {
	if promotion, err := repository.GetPromotingById(id); err != nil {
		return dto.PromotionResponseDto{}, err
	} else {
		return dto.PromotionResponseDto{
			Id:             promotion.Id,
			OriginalId:     promotion.OriginalId,
			Price:          promotion.Price,
			ExpirationDate: promotion.ExpirationDate,
		}, nil
	}
}

// TOOD: Transactional
func UploadCSV(file multipart.File) {
	go func() {
		err := truncateAndUpload(file)
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Printf("The file has been uploaded successfully")
		}
	}()
}

func truncateAndUpload(file multipart.File) error {
	defer file.Close()
	repository.DeleteAll()
	return uploadPromotionsCsv(file)
}

func uploadPromotionsCsv(file multipart.File) error {
	csvReader := csv.NewReader(file)

	for {
		row, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
			return err
		}

		log.Printf("readed: %s", row)

		originalId := uuid.Must(uuid.Parse(row[0]))
		price, _ := strconv.ParseFloat(row[1], 64)
		expirationDate, _ := time.Parse("2006-01-02 15:04:05 Z0700 MST", row[2])

		fromRowPromotion := model.Promotion{
			OriginalId:     originalId,
			Price:          price,
			ExpirationDate: expirationDate,
		}

		err = repository.Save(fromRowPromotion)

		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteAllPromotions() {
	repository.DeleteAll()
}
