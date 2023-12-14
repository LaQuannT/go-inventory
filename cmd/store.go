package cmd

import (
	"log"

	"github.com/LaQuannT/go-inventory/controller"
	"github.com/LaQuannT/go-inventory/database"
	"github.com/LaQuannT/go-inventory/services"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Stores an item",
	Long:  `Prompts user for data to create and store an item`,
	Run: func(cmd *cobra.Command, args []string) {
		dbPool, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}

		defer dbPool.Close()

		s := services.New(dbPool)

		controller.Add(s)
	},
}
