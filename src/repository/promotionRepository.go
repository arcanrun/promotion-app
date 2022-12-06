package repository

import (
	"github.com/promotionApp/src/config"
	"github.com/promotionApp/src/model"
	"log"
)

const rowAffectedTmpl = "rows affected: %d"

func GetAllPromotions() []model.Promotion {
	var promotions []model.Promotion
	result := config.DB.Debug().Find(&promotions)

	log.Printf(rowAffectedTmpl, result.RowsAffected)

	return promotions
}

func Save(promotion model.Promotion) error {
	result := config.DB.Debug().Create(&promotion)

	return result.Error
}

func GetPromotingById(id int64) (model.Promotion, error) {
	var promotion model.Promotion
	result := config.DB.Debug().First(&promotion, id)

	log.Printf(rowAffectedTmpl, result.RowsAffected)

	return promotion, result.Error
}

func DeleteAll() {
	result := config.DB.Debug().Exec("DELETE FROM promotions")

	log.Printf(rowAffectedTmpl, result.RowsAffected)
}
