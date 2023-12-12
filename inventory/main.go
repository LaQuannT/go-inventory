package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LaQuannT/inventory-mamagment-system/internal/controller"
	db "github.com/LaQuannT/inventory-mamagment-system/internal/database"
	"github.com/LaQuannT/inventory-mamagment-system/internal/services"
	"github.com/alexflint/go-arg"
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

var args struct {
	Item     bool `arg:"-i, --item" help:"add new item"`
	Category bool `arg:"-c, --category" help:"search items by category"`
	DelItem  bool `arg:"-d, --deleteItem" help:"remove item from inventory"`
	Edit     bool `arg:"-e, --edit" help:"edit item data"`
	Name     bool `arg:"-n, --name" help:"search items by name"`
	Brand    bool `arg:"-b, --brand" help:"search items by brand"`
	Search   bool `arg:"-s, --search" help:"search item by SKU code"`
}

func main() {
	arg.MustParse(&args)

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
		os.Getenv("PG_SSLMODE"),
	}

	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", env.Username, env.Password, env.Server, env.Port, env.Database, env.Ssl)

	pool, err := db.Connect(dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	s := services.New(pool)

	switch {
	case args.Brand:
		controller.BrandSearch(s)

	case args.Category:
		controller.CategorySearch(s)

	case args.DelItem:
		controller.Delete(s)

	case args.Edit:
		controller.Update(s)

	case args.Item:
		controller.Add(s)

	case args.Name:
		controller.NameSearch(s)

	case args.Search:
		controller.SearchSku(s)

	default:
		fmt.Println("Please specify a flag to perform a service. For help use -h or --help.")
		os.Exit(0)
	}
}
