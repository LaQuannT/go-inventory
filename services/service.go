package services

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Service struct {
	Db *sql.DB
}

type Item struct {
	Name, Brand, Sku, Location, Category string
}

func (s *Service) AddItem(i *Item) error {
	_, err := s.Db.Exec(`
    INSERT INTO item
    (name, brand, sku, location, category)
    values ($1, $2, $3, $4, $5)`,
		i.Name, i.Brand, i.Sku, i.Location, i.Category)
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			if e.Code == "23505" {
				// unique constraint violation branch
				return fmt.Errorf("SKU code [%s] already in use", i.Sku)
			}
		}
		return fmt.Errorf("services/AddItem: %w", err)
	}

	return nil
}

func (s *Service) SearchByCategory(category string) ([]*Item, error) {
	var res []*Item

	rows, err := s.Db.Query(`SELECT * FROM item WHERE category = $1`, category)
	if err != nil {
		return nil, fmt.Errorf("services/SearchByCategory: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		item, _, err := fromRowToItem(rows)
		if err != nil {
			return nil, fmt.Errorf("services/SearchByCategory: %w", err)
		}
		res = append(res, item)
	}
	return res, nil
}

func (s *Service) SearchByName(name string) ([]*Item, error) {
	var res []*Item

	rows, err := s.Db.Query(`SELECT * FROM item WHERE name = $1`, name)
	if err != nil {
		return nil, fmt.Errorf("services/SearchByName: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		item, _, err := fromRowToItem(rows)
		if err != nil {
			return nil, fmt.Errorf("services/SearchByName: %w", err)
		}
		res = append(res, item)
	}
	return res, nil
}

func (s *Service) SearchByBrand(brand string) ([]*Item, error) {
	var res []*Item

	rows, err := s.Db.Query(`SELECT * FROM item WHERE brand = $1`, brand)
	if err != nil {
		return nil, fmt.Errorf("services/SearchByBrand: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		item, _, err := fromRowToItem(rows)
		if err != nil {
			return nil, fmt.Errorf("services/SearchByBrand: %w", err)
		}
		res = append(res, item)
	}
	return res, nil
}

func (s *Service) SearchBySKU(sku string) (*Item, int, error) {
	var (
		res *Item
		id  int
	)

	rows, err := s.Db.Query(`SELECT * FROM item WHERE sku = $1`, sku)
	if err != nil {
		return nil, 0, fmt.Errorf("services/SearchBySKU: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		res, id, err = fromRowToItem(rows)
		if err != nil {
			return nil, 0, fmt.Errorf("services/SearchBySKU: %w", err)
		}
	}
	return res, id, nil
}

func (s *Service) DeleteItem(sku string) error {
	_, err := s.Db.Exec(`DELETE FROM item WHERE sku = $1`, sku)
	if err != nil {
		return fmt.Errorf("services/DeleteItem: %w", err)
	}
	return nil
}

func (s *Service) EditItem(i *Item, id int) error {
	_, err := s.Db.Exec(`
    UPDATE item
    SET name = $1, brand = $2, sku = $3, location = $4, category =$5
    WHERE id = $6`, i.Name, i.Brand, i.Sku, i.Location, i.Category, id)
	if err != nil {
		return fmt.Errorf("services/EditItem: %w", err)
	}

	return nil
}

func fromRowToItem(row *sql.Rows) (*Item, int, error) {
	var sku, name, brand, category, location string
	var id int

	if err := row.Scan(&id, &sku, &name, &brand, &category, &location); err != nil {
		return nil, 0, fmt.Errorf("fromRowToItem: %w", err)
	}
	return &Item{
		name,
		brand,
		sku,
		location,
		category,
	}, id, nil
}

func New(db *sql.DB) *Service {
	return &Service{Db: db}
}
