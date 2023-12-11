package controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/LaQuannT/inventory-mamagment-system/internal/services"
	"github.com/LaQuannT/inventory-mamagment-system/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const initAttempt = 1

func Add(s *services.Service) {
	var i services.Item

	i.Name = utils.ValidateInput("Name", initAttempt)
	i.Brand = utils.ValidateInput("Brand", initAttempt)
	i.Sku = utils.ValidateInput("Stock keeping unit(sku)", initAttempt)
	i.Location = utils.ValidateInput("Location", initAttempt)
	i.Category = utils.ValidateInput("Category", initAttempt)

	if err := s.AddItem(&i); err != nil {
		log.Fatal(err)
	}

	displayData(&i)
}

func CategorySearch(s *services.Service) {
	c := utils.ValidateInput("Search Category", initAttempt)

	items, err := s.SearchByCategory(c)
	if err != nil {
		log.Fatal(err)
	}

	handleMultiItemDisplay(items)
}

func NameSearch(s *services.Service) {
	n := utils.ValidateInput("Search Name", initAttempt)

	items, err := s.SearchByName(n)
	if err != nil {
		log.Fatal(err)
	}

	handleMultiItemDisplay(items)
}

func BrandSearch(s *services.Service) {
	b := utils.ValidateInput("Search Brand", initAttempt)

	items, err := s.SearchByBrand(b)
	if err != nil {
		log.Fatal(err)
	}

	handleMultiItemDisplay(items)
}

func SearchSku(s *services.Service) {
	code := utils.ValidateInput("Search SKU code", initAttempt)

	item, _, err := s.SearchBySKU(code)
	if err != nil {
		log.Fatal(err)
	}

	displayData(item)
}

func Delete(s *services.Service) {
	code := utils.ValidateInput("SKU of item to delete", initAttempt)

	if err := s.DeleteItem(code); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Item deleted: %s", code)
}

func Update(s *services.Service) {
	var updatedItem services.Item
	code := utils.ValidateInput("SKU of item to update", initAttempt)

	item, id, err := s.SearchBySKU(code)
	if err != nil {
		log.Fatal(err)
	}

	updatedItem.Name = utils.ValidateInputChange("Name", item.Name)
	updatedItem.Brand = utils.ValidateInputChange("Brand", item.Brand)
	updatedItem.Sku = utils.ValidateInputChange("SKU", item.Sku)
	updatedItem.Category = utils.ValidateInputChange("Category", item.Category)
	updatedItem.Location = utils.ValidateInputChange("Location", item.Location)

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
