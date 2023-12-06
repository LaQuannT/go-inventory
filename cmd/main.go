package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/LaQuannT/inventory-mamagment-system/internal/database"
	"github.com/joho/godotenv"
)

type env struct {
	Username string
	Password string
	Server   string
	Port     string
	Database string
	Ssl      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file...")
	}

	env := env{
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_SERVER"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("SSLMODE"),
	}

	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", env.Username, env.Password, env.Server, env.Port, env.Database, env.Ssl)

	_, err = db.Connect(dbConnStr)
	if err != nil {
		log.Fatal("Couldn't connect to database...")
	}
}
