package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBDriver   string
	AppPort    string
}

const projectDirName = "url-shortener-go" // change to relevant project name

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() *Config {
	loadEnv()

	return &Config{
		AppPort:    ":" + os.Getenv("APP_PORT"),
		DBUsername: os.Getenv("APP_DB_USERNAME"),
		DBPassword: os.Getenv("APP_DB_PASSWORD"),
		DBName:     os.Getenv("APP_DB_NAME"),
		DBDriver:   os.Getenv("APP_DB_DRIVER"),
	}
}
