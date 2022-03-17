package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"gin_and_tests/models"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dbstring := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err = gorm.Open(postgres.Open(dbstring))
	if err != nil {
		log.Panic("Error to connect to database")
	}
	DB.AutoMigrate(&models.Student{})
}
