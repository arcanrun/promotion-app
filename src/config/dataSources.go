package config

import (
	"fmt"
	"github.com/magiconair/properties"
	"github.com/promotionApp/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DataSource() *gorm.DB {
	props := properties.MustLoadFile("src/resources/application-${PROFILE}.properties", properties.UTF8)
	dbUser, _ := props.Get("datasource.user")
	dbPasswd, _ := props.Get("datasource.password")
	dbPort, _ := props.Get("datasource.port")
	dbName, _ := props.Get("datasource.name")
	dbDriver, _ := props.Get("datasource.driver")
	dbUrl, _ := props.Get("datasource.url")
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbDriver, dbUser, dbPasswd, dbUrl, dbPort, dbName)

	log.Println(dbURL)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	migrationError := db.AutoMigrate(&model.Promotion{})

	if migrationError != nil {
		log.Fatalln(migrationError.Error())
	}

	return db
}
