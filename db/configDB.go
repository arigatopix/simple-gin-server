package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})

	return db
}
