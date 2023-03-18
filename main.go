package main

import (
	"log"

	"github.com/berryfl/alb-api-ng/instance"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=postgres password=23456 dbname=alb_db port=5432"
)

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect_database_failed: %v\n", err)
	}

	inst := &instance.Instance{
		Name:        "alb-berry",
		EnableHTTP:  true,
		EnableHTTPS: false,
		UpdatedBy:   "berry",
	}
	if err := inst.Create(db); err != nil {
		log.Fatalln("create_instance_failed: exit")
	}
}
