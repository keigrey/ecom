package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost 			   string
	Port       			   string
	DBPort     			   string
	DBUser     			   string
	DBPassword 			   string
	DBName     			   string
	JWTExpirationInSeconds int
	JWTSecret			   string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		PublicHost: 			os.Getenv("DB_HOST"),
		Port:       			os.Getenv("PORT"),
		DBPort:     			os.Getenv("DB_PORT"),
		DBUser:     			os.Getenv("DB_USER"),
		DBPassword: 			os.Getenv("DB_PASSWORD"),
		DBName:     			os.Getenv("DB_NAME"),
		JWTExpirationInSeconds: 3600 * 24 * 7,
		JWTSecret: 				os.Getenv("JWT_SECRET"),
	}
}