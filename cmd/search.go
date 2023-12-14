package cmd

import (
	"log"

	"github.com/LaQuannT/go-inventory/controller"
	"github.com/LaQuannT/go-inventory/database"
	"github.com/LaQuannT/go-inventory/services"
	"github.com/spf13/cobra"
)

var (
	name     bool
	category bool
	brand    bool
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Find items",
	Long:  "Allows users to find one or several items in various ways",
	Run: func(cmd *cobra.Command, args []string) {
		dbPool, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}

		defer dbPool.Close()

		s := services.New(dbPool)

		switch {
		case name:
			controller.NameSearch(s)
		case category:
			controller.CategorySearch(s)
		case brand:
			controller.BrandSearch(s)
		default:
			controller.SearchSku(s)
		}
	},
}

func init() {
	searchCmd.Flags().BoolVarP(&name, "name", "n", false, "Specifies search by name")
	searchCmd.Flags().BoolVarP(&category, "category", "c", false, "Specifies search by category")
	searchCmd.Flags().BoolVarP(&brand, "brand", "b", false, "Specifies search by brand")
}
