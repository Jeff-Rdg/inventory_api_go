package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	InventoryProductError   = errors.New("invalid product to inventory")
	InventoryStorageError   = errors.New("invalid storage to inventory")
	InventoryProviderError  = errors.New("invalid provider to inventory")
	InventoryUnitValueError = errors.New("invalid Unit value to inventory")
	InventoryQuantityError  = errors.New("invalid Quantity to inventory")
)

type Inventory struct {
	base
	product   Product
	storage   Storage
	provider  Provider
	unitValue float64
	quantity  int
}

func NewInventory(product Product, storage Storage, provider Provider, unitValue float64, quantity int) (Inventory, error) {
	if (product == Product{}) {
		return Inventory{}, InventoryProductError
	}
	if (storage == Storage{}) {
		return Inventory{}, InventoryStorageError
	}
	if (provider == Provider{}) {
		return Inventory{}, InventoryProviderError
	}
	if unitValue <= 0 {
		return Inventory{}, InventoryUnitValueError
	}
	if quantity <= 0 {
		return Inventory{}, InventoryQuantityError
	}

	return Inventory{
		base: base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		product:   product,
		storage:   storage,
		provider:  provider,
		unitValue: unitValue,
		quantity:  quantity,
	}, nil
}
