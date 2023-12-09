package services

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/LaQuannT/inventory-mamagment-system/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Service struct {
	Db *sql.DB
}

type item struct {
	Name, Brand, Sku, Location, Category string
}

func (s *Service) AddItem() {
	defer s.Db.Close()
	var i item
	i.Name = utils.StringPrompt("Item name")
	i.Brand = utils.StringPrompt("Item brand")
	i.Sku = utils.StringPrompt("Stock keeping unit(sku)")
	i.Location = utils.StringPrompt("Location in warehouse")
	i.Category = utils.StringPrompt("Item category")

	_, err := s.Db.Exec(`
    INSERT INTO item
    (name, brand, sku, location, category)
    values ($1, $2, $3, $4, $5)`,
		i.Name, i.Brand, i.Sku, i.Location, i.Category)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("item added: %s", i.Sku)
}

func (s *Service) SearchByCategory() {
	defer s.Db.Close()

	c := utils.StringPrompt("Search category")

	rows, err := s.Db.Query(`SELECT * FROM item WHERE category = $1`, c)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totalItems := 0

	for rows.Next() {
		item, _, err := fromRowToItem(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		totalItems += 1
		displayData(item)
	}
	fmt.Println("items found:", totalItems)
}

func (s *Service) SearchByName() {
	defer s.Db.Close()

	n := utils.StringPrompt("Search name")

	rows, err := s.Db.Query(`SELECT * FROM item WHERE name = $1`, n)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totalItems := 0

	for rows.Next() {
		item, _, err := fromRowToItem(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		totalItems += 1
		displayData(item)
	}
	fmt.Println("items found:", totalItems)
}

func (s *Service) SearchByBrand() {
	defer s.Db.Close()

	b := utils.StringPrompt("Search brand")

	rows, err := s.Db.Query(`SELECT * FROM item WHERE brand = $1`, b)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totalItems := 0

	for rows.Next() {
		item, _, err := fromRowToItem(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		totalItems += 1
		displayData(item)
	}
	fmt.Println("items found:", totalItems)
}

func (s *Service) SearchBySKU() {
	defer s.Db.Close()

	sku := utils.StringPrompt("Search SKU code")

	rows, err := s.Db.Query(`SELECT * FROM item WHERE sku = $1`, sku)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totalItems := 0

	for rows.Next() {
		item, _, err := fromRowToItem(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		totalItems += 1
		displayData(item)
	}
	fmt.Println("items found:", totalItems)
}

func (s *Service) DeleteItem() {
	defer s.Db.Close()

	sku := utils.StringPrompt("Item SKU code to delete")
	_, err := s.Db.Exec(`DELETE FROM item WHERE sku = $1`, sku)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Item %s deleted...", sku)
}

func (s *Service) EditItem() {
	var i *item
	var id int
	sku := utils.StringPrompt("SKU code of item to edit")

	rows, err := s.Db.Query(`SELECT * FROM item WHERE sku = $1`, sku)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		item, itemId, err := fromRowToItem(rows)
		if err != nil {
			log.Fatal(err)
		}
		i = item
		id = itemId
	}
	name := utils.StringPrompt(fmt.Sprintf("Name (%s)", i.Name))
	brand := utils.StringPrompt(fmt.Sprintf("Brand (%s)", i.Brand))
	sku = utils.StringPrompt(fmt.Sprintf("SKU (%s)", i.Sku))
	location := utils.StringPrompt(fmt.Sprintf("Location (%s)", i.Location))
	category := utils.StringPrompt(fmt.Sprintf("Category (%s)", i.Category))

	i.Name = checkVarForChange(name, i.Name)
	i.Brand = checkVarForChange(brand, i.Brand)
	i.Sku = checkVarForChange(sku, i.Sku)
	i.Location = checkVarForChange(location, i.Location)
	i.Category = checkVarForChange(category, i.Category)

	_, err = s.Db.Exec(`
    UPDATE item
    SET name = $1, brand = $2, sku = $3, location = $4, category =$5
    WHERE id = $6`, i.Name, i.Brand, i.Sku, i.Location, i.Category, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Item updated...")
	displayData(i)
}

func fromRowToItem(row *sql.Rows) (*item, int, error) {
	var sku, name, brand, category, location string
	var id int

	if err := row.Scan(&id, &sku, &name, &brand, &category, &location); err != nil {
		return nil, 0, err
	}
	return &item{
		name,
		brand,
		sku,
		location,
		category,
	}, id, nil
}

func checkVarForChange(x, y string) string {
	if x == "" {
		return y
	}
	return x
}

func displayData(i *item) {
	toTitle := cases.Title(language.English, cases.NoLower)

	i.Name = toTitle.String(i.Name)
	i.Brand = toTitle.String(i.Brand)
	i.Sku = strings.ToUpper(i.Sku)
	i.Location = strings.ToUpper(i.Location)
	i.Category = toTitle.String(i.Category)

	fmt.Printf("[%s] Item: %s | Brand: %s | Category: %s | location: %s\n", i.Sku, i.Name, i.Brand, i.Category, i.Location)
}

func New(db *sql.DB) *Service {
	return &Service{Db: db}
}
