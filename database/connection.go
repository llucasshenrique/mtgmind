package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateConnection(path string) *gorm.DB {
	database, error := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	}
	return database
}
