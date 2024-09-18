package main

import (
	"fmt"
	"log"

	"github.com/keigrey/ecom/cmd/api"
	"github.com/keigrey/ecom/config"
	"github.com/keigrey/ecom/db"
)

func main() {
	cfg := db.PostgresConfig{
		Host:       config.Envs.PublicHost,
		DBPort:     config.Envs.DBPort,
		DBUser:     config.Envs.DBUser,
		DBPassword: config.Envs.DBPassword,
		DBName:     config.Envs.DBName,
	}

	db, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}