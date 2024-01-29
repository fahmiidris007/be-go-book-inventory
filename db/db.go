package db

import (
	"book-inventory/models"
	_ "database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connect := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}
	if db.Find(&data).RecordNotFound() {
		seederBooks(db)
	}
}

func seederBooks(db *gorm.DB) {
	data := []models.Books{
		{
			Title:       "The Lord of The Rings",
			Author:      "J.R.R. Tolkien",
			Description: "The Lord of the Rings is an epic high-fantasy novel written by English author and scholar J. R. R. Tolkien.",
			Stock:       10,
		},
		{
			Title:       "Harry Potter",
			Author:      "J.K. Rowling",
			Description: "Harry Potter is a series of seven fantasy novels written by British author J. K. Rowling.",
			Stock:       10,
		},
		{
			Title:       "The Hobbit",
			Author:      "J.R.R. Tolkien",
			Description: "The Hobbit, or There and Back Again is a children's fantasy novel by English author J. R. R. Tolkien.",
			Stock:       10,
		},
		{
			Title:       "The Little Prince",
			Author:      "Antoine de Saint-Exupéry",
			Description: "The Little Prince is a novella by French aristocrat, writer, and aviator Antoine de Saint-Exupéry.",
			Stock:       10,
		},
	}

	for _, v := range data {
		db.Create(&v)
	}
}
