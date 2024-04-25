package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDBPostgres() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
