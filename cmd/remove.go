package cmd

import (
	"log"

	"github.com/LaQuannT/go-inventory/controller"
	"github.com/LaQuannT/go-inventory/database"
	"github.com/LaQuannT/go-inventory/services"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an item",
	Long:  `Prompts user for stock keepin unit(SKU) and removes item if found`,
	Run: func(cmd *cobra.Command, args []string) {
		dbPool, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}

		defer dbPool.Close()

		s := services.New(dbPool)

		controller.Delete(s)
	},
}
