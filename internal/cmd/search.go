package cmd

import (
	"context"
	"log"

	"github.com/LaQuannT/go-inventory/internal/database"
	"github.com/LaQuannT/go-inventory/internal/database/repository"
	"github.com/LaQuannT/go-inventory/internal/service"
	"github.com/spf13/cobra"
)

var (
	name     bool
	category bool
	brand    bool
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for one or more items",
	Long:  "Search for one or more items by Stock Keeping Unit(default), Name, Brand, or Category",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close(context.Background())

		r := repository.New(conn)
		s := service.New(r)

		switch {
		case name:
			s.NameSearch()
		case category:
			s.CategorySearch()
		case brand:
			s.BrandSearch()
		default:
			s.SKUSearch()
		}
	},
}

func init() {
	searchCmd.Flags().BoolVarP(&name, "name", "n", false, "Specifies search by name")
	searchCmd.Flags().BoolVarP(&category, "category", "c", false, "Specifies search by category")
	searchCmd.Flags().BoolVarP(&brand, "brand", "b", false, "Specifies search by brand")
}
