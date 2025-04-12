package db

import (
	"log"

	"github.com/XoDeR/simple-rest-api-go/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// TODO: move to env var
	dbURL := "postgres://postgres@localhost:5432/go-crud-01"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Database connection established successfully!")

	db.AutoMigrate(&models.Book{})

	return db
}
