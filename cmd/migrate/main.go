package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // for file source
	_ "github.com/jackc/pgx/v5/stdlib"                   // for pgx postgres driver
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

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)] 
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	log.Println("Migrations applied successfully!")
}