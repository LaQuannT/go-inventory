package controller

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/LaQuannT/go-inventory/services"
	"github.com/LaQuannT/go-inventory/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const initAttempt = 1

func Add(s *services.Service) {
	name, err := utils.ValidateInputData(os.Stdin, "Name", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	brand, err := utils.ValidateInputData(os.Stdin, "Brand", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	sku, err := utils.ValidateInputData(os.Stdin, "Stock keeping unit(sku)", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	category, err := utils.ValidateInputData(os.Stdin, "Category", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	location, err := utils.ValidateInputData(os.Stdin, "Location", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	i := services.Item{
		Name:     name,
		Brand:    brand,
		Sku:      sku,
		Category: category,
		Location: location,
	}

	if err := s.AddItem(&i); err != nil {
		log.Fatal(err)
	}

	displayData(&i)
}

func CategorySearch(s *services.Service) {
	c, err := utils.ValidateInputData(os.Stdin, "Search Category", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	items, err := s.SearchByCategory(c)
	if err != nil {
		log.Fatal(err)
	}

	handleMultiItemDisplay(items)
}

func NameSearch(s *services.Service) {
	n, err := utils.ValidateInputData(os.Stdin, "Search Name", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	items, err := s.SearchByName(n)
	if err != nil {
		log.Fatal(err)
	}

	handleMultiItemDisplay(items)
}

func BrandSearch(s *services.Service) {
	b, err := utils.ValidateInputData(os.Stdin, "Search Brand", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	items, err := s.SearchByBrand(b)
	if err != nil {
		log.Fatal(err)
	}

	handleMultiItemDisplay(items)
}

func SearchSku(s *services.Service) {
	code, err := utils.ValidateInputData(os.Stdin, "Search SKU code", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	item, _, err := s.SearchBySKU(code)
	if err != nil {
		log.Fatal(err)
	}

	displayData(item)
}

func Delete(s *services.Service) {
	code, err := utils.ValidateInputData(os.Stdin, "SKU of item to delete", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.DeleteItem(code); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Item deleted: %s", code)
}

func Update(s *services.Service) {
	code, err := utils.ValidateInputData(os.Stdin, "SKU of item to update", initAttempt)
	if err != nil {
		log.Fatal(err)
	}

	item, id, err := s.SearchBySKU(code)
	if err != nil {
		log.Fatal(err)
	}

	name, err := utils.ValidateInputDataChange(os.Stdin, "Name", item.Name)
	if err != nil {
		log.Fatal(err)
	}

	brand, err := utils.ValidateInputDataChange(os.Stdin, "Brand", item.Brand)
	if err != nil {
		log.Fatal(err)
	}

	sku, err := utils.ValidateInputDataChange(os.Stdin, "SKU", item.Sku)
	if err != nil {
		log.Fatal(err)
	}

	category, err := utils.ValidateInputDataChange(os.Stdin, "Category", item.Category)
	if err != nil {
		log.Fatal(err)
	}

	location, err := utils.ValidateInputDataChange(os.Stdin, "Location", item.Location)
	if err != nil {
		log.Fatal(err)
	}

	updatedItem := services.Item{
		Name:     name,
		Brand:    brand,
		Sku:      sku,
		Category: category,
		Location: location,
	}

	if err := s.EditItem(&updatedItem, id); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Item update:")
	displayData(&updatedItem)
}

func displayData(i *services.Item) {
	toTitle := cases.Title(language.English, cases.NoLower)

	i.Name = toTitle.String(i.Name)
	i.Brand = toTitle.String(i.Brand)
	i.Sku = strings.ToUpper(i.Sku)
	i.Location = strings.ToUpper(i.Location)
	i.Category = toTitle.String(i.Category)

	fmt.Printf("[%s] Item: %s | Brand: %s | Category: %s | location: %s\n", i.Sku, i.Name, i.Brand, i.Category, i.Location)
}

func handleMultiItemDisplay(items []*services.Item) {
	for _, Item := range items {
		displayData(Item)
	}
	fmt.Printf("total items found: %d\n", len(items))
}
