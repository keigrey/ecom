package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func NewPostgresStorage(cfg PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
						cfg.DBUser, 
						cfg.DBPassword, 
						cfg.Host, 
						cfg.DBPort, 
						cfg.DBName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")

	return db, nil
}
