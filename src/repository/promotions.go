package repository

import (
	"github.com/google/uuid"
	"github.com/promotionApp/src/model"
	"math/rand"
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
