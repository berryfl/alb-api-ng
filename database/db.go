package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=postgres password=23456 dbname=alb_db port=5432"
)

var (
	db *gorm.DB
)

func InitDB() error {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("connect_database_failed: %v\n", err)
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
