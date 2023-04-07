package main

import (
	"log"

	"github.com/berryfl/alb-api-ng/database"
	"github.com/berryfl/alb-api-ng/web"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("alb")
	viper.AddConfigPath("config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read_config_failed: %v\n", err)
	}

	connectParams := &database.ConnectParams{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
		Port:     viper.GetUint16("database.port"),
	}

	if err := database.InitDB(connectParams); err != nil {
		log.Fatalln("initialize_database_failed: exit")
	}

	r := web.NewRouter()
	if err := r.Run(":18080"); err != nil {
		log.Fatalf("run_router_failed: %v\n", err)
	}
}
