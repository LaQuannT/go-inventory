package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LaQuannT/go-inventory/internal/model"
	"github.com/jackc/pgx/v5"
)

type itemRepository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *itemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(i *model.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `
  INSERT INTO item (name, brand, stock_keeping_unit, category, location, amount)
  VALUES ($1, $2, $3, $4, $5, $6);
  `
	_, err := r.db.Exec(ctx, stmt, i.Name, i.Brand, i.Sku, i.Category, i.Location)
	if err != nil {
		return fmt.Errorf("unable to add item: %w", err)
	}
	return nil
}

func (r *itemRepository) SearchCategory(c string) ([]*model.Item, error) {
	items := make([]*model.Item, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `SELECT * FROM item WHERE category=$1;`

	rows, err := r.db.Query(ctx, stmt, c)
	if err != nil {
		return nil, fmt.Errorf("unable to find items in category %q: %w", c, err)
	}
	defer rows.Close()

	for rows.Next() {
		i, err := rowToItem(rows)
		if err != nil {
			return nil, fmt.Errorf("unable to list item data: %w", err)
		}
		items = append(items, i)
	}

	return items, nil
}

func (r *itemRepository) SearchName(n string) ([]*model.Item, error) {
	items := make([]*model.Item, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `SELECT * FROM item WHERE name=$1;`

	rows, err := r.db.Query(ctx, stmt, n)
	if err != nil {
		return nil, fmt.Errorf("unable to find items with name %q: %w", n, err)
	}
	defer rows.Close()

	for rows.Next() {
		i, err := rowToItem(rows)
		if err != nil {
			return nil, fmt.Errorf("unable to list item data: %w", err)
		}
		items = append(items, i)
	}

	return items, nil
}

func (r *itemRepository) SearchBrand(b string) ([]*model.Item, error) {
	items := make([]*model.Item, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `SELECT * FROM item WHERE brand=$1;`

	rows, err := r.db.Query(ctx, stmt, b)
	if err != nil {
		return nil, fmt.Errorf("unable to find items from brand %q: %w", b, err)
	}
	defer rows.Close()

	for rows.Next() {
		i, err := rowToItem(rows)
		if err != nil {
			return nil, fmt.Errorf("unable to list item data: %w", err)
		}
		items = append(items, i)
	}

	return items, nil
}

func (r *itemRepository) SearchSKU(sku string) (*model.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `SELECT * FROM item WHERE stock_keeping_unit=$1;`

	rows, err := r.db.Query(ctx, stmt, sku)
	if err != nil {
		return nil, fmt.Errorf("unable to find items with stock keeping unit %q: %w", sku, err)
	}
	defer rows.Close()

	for rows.Next() {
		return rowToItem(rows)
	}

	return nil, nil
}

func (r *itemRepository) Delete(sku string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `DELETE FROM item WHERE sku=$1;`
	_, err := r.db.Exec(ctx, stmt, sku)
	if err != nil {
		return fmt.Errorf("unable to delete item: %w", err)
	}

	return nil
}

func (r *itemRepository) Update(i *model.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `UPDATE item SET name=$1, brand=$2, stock_keeping_unit=$3, category=$4, location=$5,
  amount=$6 WHERE id=$7;
  `
	_, err := r.db.Exec(ctx, stmt, i.Name, i.Brand, i.Sku, i.Category, i.Location, i.Amount, i.ID)
	if err != nil {
		return fmt.Errorf("unable to update item: %w", err)
	}

	return nil
}

func rowToItem(r pgx.Rows) (*model.Item, error) {
	i := new(model.Item)

	if err := r.Scan(i.ID, i.Name, i.Brand, i.Sku, i.Category, i.Location, i.Amount); err != nil {
		return nil, errors.New("pgx pq row to struct conversion error")
	}
	return i, nil
}
