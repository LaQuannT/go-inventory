package service

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/LaQuannT/go-inventory/internal/model"
)

var prompts = [6]string{"Name", "Brand", "Stock Keeping Unit(SKU)", "Category", "Location", "Amount"}

type service struct {
	repository model.ItemRepository
}

func (s *service) Create() {
	var name, brand, sku, category, location string

	inputs := []string{name, brand, sku, category, location}

	for i := 0; i < len(prompts)-2; i++ {
		input, err := validateInput(os.Stdin, prompts[i])
		if err != nil {
			log.Fatal(err)
		}
		inputs[i] = input
	}

	n, err := validateInput(os.Stdin, prompts[5])
	if err != nil {
		log.Fatal(err)
	}

	amount, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("%q must be a valid number", prompts[5])
	}

	i := &model.Item{Name: name, Brand: brand, Sku: sku, Category: category, Location: location, Amount: amount}

	if err = s.repository.Create(i); err != nil {
		log.Fatal(err)
	}

	displayData(i)
}

func (s service) Add() {
	sku, err := validateInput(os.Stdin, prompts[2])
	if err != nil {
		log.Fatal(err)
	}

	n, err := validateInput(os.Stdin, prompts[5])
	if err != nil {
		log.Fatal(err)
	}

	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("%q must be a valid number", prompts[5])
	}

	i, err := s.repository.SearchSKU(sku)
	if err != nil {
		log.Fatalf("unable to find item with SKU %q: %v", sku, err)
	}

	i.Amount += num

	if err = s.repository.Update(i); err != nil {
		log.Fatalf("unable to add to item amount: %v", err)
	}
	displayData(i)
}

func (s *service) Subtract() {
	sku, err := validateInput(os.Stdin, prompts[2])
	if err != nil {
		log.Fatal(err)
	}

	n, err := validateInput(os.Stdin, prompts[5])
	if err != nil {
		log.Fatal(err)
	}

	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("%q must be a valid number", prompts[5])
	}

	i, err := s.repository.SearchSKU(sku)
	if err != nil {
		log.Fatalf("unable to find item with SKU %q: %v", sku, err)
	}

	i.Amount -= num

	if err = s.repository.Update(i); err != nil {
		log.Fatalf("unable to remove item amount: %v", err)
	}

	displayData(i)
}

func (s *service) NameSearch() {
	n, err := validateInput(os.Stdin, prompts[0])
	if err != nil {
		log.Fatal(err)
	}

	items, err := s.repository.SearchName(n)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range items {
		displayData(i)
	}
}

func (s *service) CategorySearch() {
	c, err := validateInput(os.Stdin, prompts[3])
	if err != nil {
		log.Fatal(err)
	}

	items, err := s.repository.SearchCategory(c)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range items {
		displayData(i)
	}
}

func (s *service) BrandSearch() {
	b, err := validateInput(os.Stdin, prompts[1])
	if err != nil {
		log.Fatal(err)
	}

	items, err := s.repository.SearchBrand(b)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range items {
		displayData(i)
	}
}

func (s *service) SKUSearch() {
	sku, err := validateInput(os.Stdin, prompts[2])
	if err != nil {
		log.Fatal(err)
	}

	i, err := s.repository.SearchSKU(sku)
	if err != nil {
		log.Fatal(err)
	}

	displayData(i)
}

func (s *service) Delete() {
	sku, err := validateInput(os.Stdin, prompts[2])
	if err != nil {
		log.Fatal(err)
	}

	if err = s.repository.Delete(sku); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("items deleted: %q", sku)
}

func (s *service) Update() {
	sku, err := validateInput(os.Stdin, prompts[2])
	if err != nil {
		log.Fatal(err)
	}

	i, err := s.repository.SearchSKU(sku)
	if err != nil {
		log.Fatal(err)
	}

	inputs := []string{i.Name, i.Brand, i.Sku, i.Category, i.Category}

	for i := 0; i < len(prompts)-2; i++ {
		prompt := fmt.Sprintf("%s [ %s ]", prompts[i], inputs[i])

		input, err := validateInput(os.Stdin, prompt)
		if err != nil {
			log.Fatal(err)
		}

		if input == "" {
			continue
		} else {
			inputs[i] = input
		}
	}

	prompt := fmt.Sprintf("%s [ %d ]", prompts[5], i.Amount)
	numStr, err := validateInput(os.Stdin, prompt)
	if err != nil {
		log.Fatal(err)
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatalf("%q must be a valid number", prompts[5])
	}

	if num != i.Amount {
		i.Amount = num
	}

	if err = s.repository.Update(i); err != nil {
		log.Fatal(err)
	}

	displayData(i)
}
