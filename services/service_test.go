package services

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var ErrDuplicateSku = errors.New("sku already in use")

func initMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	return mockDb, mock
}

func TestAddItem(t *testing.T) {
	testCases := []struct {
		desc        string
		expectErr   bool
		item        *Item
		mockHandler func(mock sqlmock.Sqlmock, i *Item)
	}{
		{
			desc: "insert-item",
			item: &Item{
				Name:     "iphone 12 pro",
				Brand:    "apple",
				Sku:      "aap12p21",
				Category: "phone",
				Location: "storage room",
			},
			mockHandler: func(mock sqlmock.Sqlmock, i *Item) {
				mock.ExpectExec("INSERT").
					WithArgs(i.Name, i.Brand, i.Sku, i.Location, i.Category).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			desc:      "duplicate-sku-insert",
			expectErr: true,
			item: &Item{
				Name:     "iphone 12",
				Brand:    "apple",
				Sku:      "aap1221",
				Category: "phone",
				Location: "storage room",
			},
			mockHandler: func(mock sqlmock.Sqlmock, i *Item) {
				mock.ExpectExec("INSERT").
					WithArgs(i.Name, i.Brand, i.Sku, i.Location, i.Category).
					WillReturnError(ErrDuplicateSku)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockDb, mock := initMock(t)
			defer mockDb.Close()

			s := &Service{Db: mockDb}

			tC.mockHandler(mock, tC.item)

			if err := s.AddItem(tC.item); err != nil {
				switch {
				case tC.expectErr:
					if !errors.Is(err, ErrDuplicateSku) {
						t.Errorf("expected error '%s', got error '%s'", ErrDuplicateSku, err)
					}
				default:
					t.Errorf("unexpected error inserting item: %s", err)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSearchByCategory(t *testing.T) {
	testCases := []struct {
		desc          string
		category      string
		expectedItems []*Item
	}{
		{
			desc:     "search-known-category",
			category: "watch",
			expectedItems: []*Item{
				{Name: "datejust 36mm", Brand: "rolex", Sku: "rlx36wg", Category: "watch", Location: "safe"},
				{Name: "santos de cartier", Brand: "cartier", Sku: "sdc1904", Category: "watch", Location: "safe"},
			},
		},
		{
			desc:          "search-unknown-category",
			category:      "gaming",
			expectedItems: []*Item{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockDb, mock := initMock(t)
			defer mockDb.Close()

			s := Service{Db: mockDb}

			rows := mock.NewRows([]string{"id", "sku", "name", "brand", "category", "location"})
			for i, row := range tC.expectedItems {
				rows.AddRow(i, row.Sku, row.Name, row.Brand, row.Category, row.Location)
			}

			mock.ExpectQuery("SELECT").WithArgs(tC.category).
				WillReturnRows(rows)

			items, err := s.SearchByCategory(tC.category)
			if err != nil {
				t.Fatalf("unexpected error searching category: %s", err)
			}

			if len(items) != len(tC.expectedItems) {
				t.Errorf("expected '%d' items, got '%d' items", len(tC.expectedItems), len(items))
			}

			for i := range items {
				if items[i].Sku != tC.expectedItems[i].Sku {
					t.Errorf("expected sku '%s', got sku '%s'", tC.expectedItems[i].Sku, items[i].Sku)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSearchByBrand(t *testing.T) {
	testCases := []struct {
		desc          string
		brand         string
		expectedItems []*Item
	}{
		{
			desc:  "search-known-brand",
			brand: "apple",
			expectedItems: []*Item{
				{Name: "iphone 15", Brand: "apple", Sku: "aap1523", Category: "phone", Location: "storage room"},
				{Name: "iphone 14 pro", Brand: "apple", Sku: "aap14p22", Category: "phone", Location: "storage room"},
			},
		},
		{
			desc:          "search-unknown-brand",
			brand:         "samsung",
			expectedItems: []*Item{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockDb, mock := initMock(t)
			defer mockDb.Close()

			s := Service{Db: mockDb}

			rows := mock.NewRows([]string{"id", "sku", "name", "brand", "category", "location"})
			for i, row := range tC.expectedItems {
				rows.AddRow(i, row.Sku, row.Name, row.Brand, row.Category, row.Location)
			}

			mock.ExpectQuery("SELECT").WithArgs(tC.brand).
				WillReturnRows(rows)

			items, err := s.SearchByBrand(tC.brand)
			if err != nil {
				t.Fatalf("unexpected error searching brand: %s", err)
			}

			if len(items) != len(tC.expectedItems) {
				t.Errorf("expected '%d' items, got '%d' items", len(tC.expectedItems), len(items))
			}

			for i := range items {
				if items[i].Sku != tC.expectedItems[i].Sku {
					t.Errorf("expected sku '%s', got sku '%s'", tC.expectedItems[i].Sku, items[i].Sku)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSearchByName(t *testing.T) {
	testCases := []struct {
		desc          string
		name          string
		expectedItems []*Item
	}{
		{
			desc: "search-known-name",
			name: "playstation 5",
			expectedItems: []*Item{
				{Name: "playstation 5", Brand: "sony", Sku: "snyp5d23", Category: "gaming", Location: "game room"},
			},
		},
		{
			desc:          "search-unknown-name",
			name:          "game cube",
			expectedItems: []*Item{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockDb, mock := initMock(t)
			defer mockDb.Close()

			s := Service{Db: mockDb}

			rows := mock.NewRows([]string{"id", "sku", "name", "brand", "category", "location"})
			for i, row := range tC.expectedItems {
				rows.AddRow(i, row.Sku, row.Name, row.Brand, row.Category, row.Location)
			}

			mock.ExpectQuery("SELECT").WithArgs(tC.name).
				WillReturnRows(rows)

			items, err := s.SearchByName(tC.name)
			if err != nil {
				t.Fatalf("unexpected error searching name: %s", err)
			}

			if len(items) != len(tC.expectedItems) {
				t.Errorf("expected '%d' items, got '%d' items", len(tC.expectedItems), len(items))
			}

			for i := range items {
				if items[i].Sku != tC.expectedItems[i].Sku {
					t.Errorf("expected sku '%s', got sku '%s'", tC.expectedItems[i].Sku, items[i].Sku)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSearchBySku(t *testing.T) {
	testCases := []struct {
		desc         string
		sku          string
		expectedItem *Item
	}{
		{
			desc:         "search-known-sku",
			sku:          "lmf22r",
			expectedItem: &Item{Name: "f22-raptor", Brand: "lockheed martin", Sku: "lmf22r", Category: "jet", Location: "hanger"},
		},
		{
			desc:         "search-unknown-sku",
			sku:          "gg76h98",
			expectedItem: &Item{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockDb, mock := initMock(t)
			defer mockDb.Close()

			s := Service{Db: mockDb}

			rows := mock.NewRows([]string{"id", "sku", "name", "brand", "category", "location"}).
				AddRow(1, tC.expectedItem.Sku, tC.expectedItem.Name, tC.expectedItem.Brand, tC.expectedItem.Category, tC.expectedItem.Location)

			mock.ExpectQuery("SELECT").WithArgs(tC.sku).
				WillReturnRows(rows)

			item, _, err := s.SearchBySKU(tC.sku)
			if err != nil {
				t.Fatalf("unexpected error searching sku: %s", err)
			}

			if item.Sku != tC.expectedItem.Sku {
				t.Errorf("expected '%v', got '%v'", tC.expectedItem, item)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDeleteItem(t *testing.T) {
	mockDb, mock := initMock(t)
	defer mockDb.Close()

	s := &Service{Db: mockDb}
	sku := "aap12p21"

	mock.ExpectExec("DELETE").WithArgs(sku).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := s.DeleteItem(sku); err != nil {
		t.Errorf("unexpected error deleteing item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc        string
		updatedItem *Item
		expectErr   bool
		mockHandler func(mock sqlmock.Sqlmock, i *Item, id int)
	}{
		{
			desc:        "valid-item-update",
			updatedItem: &Item{Name: "duel sense", Brand: "sony", Sku: "cp5ds23", Category: "gaming", Location: "game room"},
			mockHandler: func(mock sqlmock.Sqlmock, i *Item, id int) {
				mock.ExpectExec("UPDATE").WithArgs(i.Name, i.Brand, i.Sku, i.Location, i.Category, id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			desc:        "duplicate-sku-update",
			updatedItem: &Item{Name: "duel shock", Brand: "sony", Sku: "cp5ds23", Category: "gaming", Location: "game room"},
			expectErr:   true,
			mockHandler: func(mock sqlmock.Sqlmock, i *Item, id int) {
				mock.ExpectExec("UPDATE").WithArgs(i.Name, i.Brand, i.Sku, i.Location, i.Category, id).
					WillReturnError(ErrDuplicateSku)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockDb, mock := initMock(t)
			defer mockDb.Close()

			s := &Service{Db: mockDb}
			id := 1

			tC.mockHandler(mock, tC.updatedItem, id)

			if err := s.EditItem(tC.updatedItem, id); err != nil {
				switch {
				case tC.expectErr:
					if !errors.Is(err, ErrDuplicateSku) {
						t.Errorf("expected error '%s', got error '%s'", ErrDuplicateSku, err)
					}
				default:
					t.Errorf("unexpected error updating item: %s", err)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
