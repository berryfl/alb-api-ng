package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type ConnectParams struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     uint16
}

func InitDB(params *ConnectParams) error {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		params.Host,
		params.User,
		params.Password,
		params.DBName,
		params.Port,
	)

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
