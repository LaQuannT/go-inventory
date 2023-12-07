package services

import (
	"database/sql"
	"log"

	"github.com/LaQuannT/inventory-mamagment-system/internal/utils"
)

type Service struct {
	Db *sql.DB
}

type item struct {
	Name, Brand, Sku, Location, Category string
	Stock                                int
}

func (s *Service) AddItem() {
	var i item
	i.Name = utils.StringPrompt("Item name")
	i.Brand = utils.StringPrompt("Item brand")
	i.Sku = utils.StringPrompt("Stock keeping unit(sku)")
	i.Location = utils.StringPrompt("Location in warehouse")
	i.Category = utils.StringPrompt("Item category")

	_, err := s.Db.Exec(`
    INSERT INTO stock 
    (name, brand, sku, location, category)
    values ($1, $2, $3, $4, $5, $6)`,
		i.Name, i.Brand, i.Sku, i.Location, i.Category)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Service) AddCategory() {
	category := utils.StringPrompt("Category name")

	_, err := s.Db.Exec(`
    INSERT INTO category (name)
    values ($1)`, category)
	if err != nil {
		log.Fatal(err)
	}
}

func New(db *sql.DB) *Service {
	return &Service{Db: db}
}
