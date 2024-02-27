package model

// Stock Keeping Unit (sku)
type Item struct {
	Name     string
	Brand    string
	Sku      string
	Category string
	Location string
	ID       int
}

type ItemRepository interface {
	Add(Item) error
	SearchCategory(string) ([]*Item, error)
	SearchName(string) ([]*Item, error)
	SearchSKU(string) ([]*Item, error)
	SearchBrand(string) ([]*Item, error)
	Delete(string) error
	Update(*Item) error
}
