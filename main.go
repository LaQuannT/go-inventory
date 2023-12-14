package main

import (
	"log"

	"github.com/LaQuannT/go-inventory/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file...")
	}

	cmd.Execute()
}
