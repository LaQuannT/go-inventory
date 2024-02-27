package main

import (
	"context"
	"log"
	"os"

	"github.com/LaQuannT/go-inventory/internal/database"
)

func main() {
	log.Println("connection to database")
	dbConn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	dbConn.Close(context.Background())
	os.Exit(0)
}
