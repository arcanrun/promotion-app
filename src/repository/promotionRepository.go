package repository

import (
	"github.com/promotionApp/src/config"
	"github.com/promotionApp/src/model"
	"log"
)

const rowAffectedTmpl = "rows affected: %d"

func GetAllPromotions() []model.Promotion {
	var promotions []model.Promotion
	result := config.DataSource().Find(&promotions)

	log.Printf(rowAffectedTmpl, result.RowsAffected)

	return promotions
}

func Save(promotion model.Promotion) error {
	result := config.DataSource().Create(&promotion)

	return result.Error
}

func GetPromotingById(id int64) (model.Promotion, error) {
	var promotion model.Promotion
	result := config.DataSource().First(&promotion, id)

	log.Printf(rowAffectedTmpl, result.RowsAffected)

	return promotion, result.Error
}
