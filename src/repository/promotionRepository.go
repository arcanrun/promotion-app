package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/promotionApp/src/model"
	"math/rand"
	"strconv"
	"time"
)

var Promotions = []model.Promotion{
	{
		Id:             1,
		OriginalId:     uuid.Must(uuid.Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")),
		Price:          rand.ExpFloat64(),
		ExpirationDate: time.Now(),
	},
	{
		Id:             2,
		OriginalId:     uuid.Must(uuid.Parse("4dca3b5a-085a-455c-8c76-9ddbf76e565a")),
		Price:          rand.ExpFloat64(),
		ExpirationDate: time.Now(),
	},
}

func GetAllPromotions() []model.Promotion {
	return Promotions
}

func Save(promotion model.Promotion) {
	Promotions = append(Promotions, promotion)
}

func GetPromotingById(id int64) (model.Promotion, error) {
	for _, promotion := range Promotions {
		if promotion.Id == id {
			return promotion, nil
		}
	}
	return model.Promotion{}, errors.New(fmt.Sprintf("Promotion with id=%s is not found",
		strconv.FormatInt(id, 10)))
}
