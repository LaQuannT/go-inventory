package cmd

import (
	"context"
	"log"

	"github.com/LaQuannT/go-inventory/internal/database"
	"github.com/LaQuannT/go-inventory/internal/database/repository"
	"github.com/LaQuannT/go-inventory/internal/service"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Stores new item",
	Long:  `Creates and stores an item of to your database`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close(context.Background())

		r := repository.New(conn)
		s := service.New(r)
		s.Create()
	},
}
