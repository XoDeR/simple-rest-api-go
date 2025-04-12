package main

import (
	"encoding/json"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"country"` // country becomes Desc
}

func main() {
	// Database connection
	dbURL := "postgres://postgres@localhost:5432/go-crud-01"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}

	log.Println("Database connected!")

	// Ensure the 'books' table exists
	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatalln("Failed to migrate 'books' model:", err)
	}

	// Read books from JSON file
	file, err := os.Open("seed/books.json")
	if err != nil {
		log.Fatalln("Failed to open books.json:", err)
	}
	defer file.Close()

	var books []Book
	err = json.NewDecoder(file).Decode(&books)
	if err != nil {
		log.Fatalln("Failed to decode JSON:", err)
	}

	// Insert books into the database
	for _, book := range books {
		err = db.Create(&book).Error
		if err != nil {
			log.Printf("Failed to insert book '%s': %v", book.Title, err)
		} else {
			log.Printf("Inserted book '%s' successfully!", book.Title)
		}
	}

	log.Println("Seeding completed!")
}
