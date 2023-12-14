package cmd

import (
	"log"

	"github.com/LaQuannT/go-inventory/controller"
	"github.com/LaQuannT/go-inventory/database"
	"github.com/LaQuannT/go-inventory/services"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "allows item updates",
	Long:  "Prompts user for stock keeping unit(SKU) of item to update, and prompts user for updated item data. Hiting enter with out entering new data will input original data",
	Run: func(cmd *cobra.Command, args []string) {
		dbPool, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}

		defer dbPool.Close()

		s := services.New(dbPool)

		controller.Update(s)
	},
}
