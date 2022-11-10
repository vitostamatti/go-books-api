package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dbSchema := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbURL := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":5432/" + dbSchema
	// dbURL := "postgres://admin:admin@0.0.0.0:5432/books"

	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
