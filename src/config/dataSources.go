package config

import (
	"github.com/promotionApp/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DataSource() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/promotion_app" //TODO: envs

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Promotion{})

	return db
}
