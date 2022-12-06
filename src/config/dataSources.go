package config

import (
	"fmt"
	"github.com/magiconair/properties"
	"github.com/promotionApp/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DataSource() {
	props := properties.MustLoadFile("src/resources/application-${PROFILE}.properties", properties.UTF8)
	dbUser, _ := props.Get("datasource.user")
	dbPasswd, _ := props.Get("datasource.password")
	dbPort, _ := props.Get("datasource.port")
	dbName, _ := props.Get("datasource.name")
	dbDriver, _ := props.Get("datasource.driver")
	dbUrl, _ := props.Get("datasource.url")
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbDriver, dbUser, dbPasswd, dbUrl, dbPort, dbName)

	log.Println(dbURL)

	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Database connection is established: %s %s", dbUrl, dbPort)
	}

	migrationError := database.AutoMigrate(&model.Promotion{})

	if migrationError != nil {
		log.Fatalln(migrationError.Error())
	}

	DB = database
}
