package service

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/LaQuannT/go-inventory/internal/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func promptForData(r io.Reader, prompt string) (string, error) {
	rdr := bufio.NewReader(r)
	fmt.Printf("%s: ", prompt)

	input, err := rdr.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("unable to process input: %w", err)
	}

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return input, nil
}

func validateInput(r io.Reader, prompt string) (string, error) {
	input, err := promptForData(r, prompt)
	if err != nil {
		return "", err
	}

	switch prompt {
	case "Name", "Brand", "Stock Keeping Unit(SKU)":
		if input == "" {
			return "", fmt.Errorf("%q must contain one or more characters", prompt)
		}
	case "Amount":
		if input == "" {
			return "", fmt.Errorf("%q must contain a valid number", prompt)
		}
	}

	return input, nil
}

func displayData(i *model.Item) {
	toTitle := cases.Title(language.English, cases.NoLower)

	i.Name = toTitle.String(i.Name)
	i.Brand = toTitle.String(i.Brand)
	i.Sku = strings.ToUpper(i.Sku)
	i.Location = strings.ToUpper(i.Location)
	i.Category = toTitle.String(i.Category)

	fmt.Printf("[%s] Name: %s | Brand: %s | Category: %s | location: %s | Stock: %d\n", i.Sku, i.Name, i.Brand, i.Category, i.Location, i.Amount)
}
