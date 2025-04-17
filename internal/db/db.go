package db

import (
	"GoApi/internal/src"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}

	if err := db.AutoMigrate(&src.Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	return db, nil
}
