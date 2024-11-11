package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"

)

type Config struct {
	HttpServer httpServer `json:"http_server"`
	DbConfig   dbConfig   `json:"db_config"`
}

type dbConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
	DbSllMode  string `env:"DB_SLL_MODE"`
	DbTimeZone string `env:"DB_TIME_ZONE"`
}

type httpServer struct {
	HttpHost string `env:"HTTP_HOST"`
	HttpPort string `env:"HTTP_PORT"`
}

func GetConfig() *Config {
	// .env faýlyny yüklemek
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconfig := dbConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbSllMode:  os.Getenv("DB_SLL_MODE"),
		DbTimeZone: os.Getenv("DB_TIME_ZONE"),
	}

	httpServerConfig := httpServer{
		HttpHost: os.Getenv("HTTP_HOST"),
		HttpPort: os.Getenv("HTTP_PORT"),
	}

	return &Config{
		DbConfig:   dbconfig,
		HttpServer: httpServerConfig,
	}
}
