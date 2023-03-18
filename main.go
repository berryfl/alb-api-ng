package main

import (
	"log"

	"github.com/berryfl/alb-api-ng/database"
	"github.com/berryfl/alb-api-ng/web"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalln("initialize_database_failed: exit")
	}

	r := web.NewRouter()
	if err := r.Run(":18080"); err != nil {
		log.Fatalf("run_router_failed: %v\n", err)
	}
}
