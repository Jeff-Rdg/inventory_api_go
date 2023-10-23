package test

import (
	"errors"
	"inventory_api/entity"
	"testing"
)

func TestInventory_NewInventory(t *testing.T) {

	type TestCase struct {
		TestName      string
		product       entity.Product
		storage       entity.Storage
		provider      entity.Provider
		unitValue     float64
		quantity      int
		ExpectedError error
	}

	address, _ := entity.NewAddress("test", "test", "test", "123", "test", "test")
	product, _ := entity.NewProduct("test", "test", "OTHER")
	storage, _ := entity.NewStorage("test", "test")
	provider, _ := entity.NewProvider(address, "test", "test")

	testCases := []TestCase{
		{
			TestName:      "Invalid Product",
			product:       entity.Product{},
			provider:      provider,
			storage:       storage,
			unitValue:     1,
			quantity:      1,
			ExpectedError: entity.InventoryProductError,
		},
		{
			TestName:      "Valid inventory",
			product:       product,
			provider:      provider,
			storage:       storage,
			unitValue:     1,
			quantity:      1,
			ExpectedError: nil,
		},
		{
			TestName:      "Invalid storage",
			product:       product,
			provider:      provider,
			storage:       entity.Storage{},
			unitValue:     1,
			quantity:      1,
			ExpectedError: entity.InventoryStorageError,
		},
		{
			TestName:      "Invalid provider",
			product:       product,
			provider:      entity.Provider{},
			storage:       storage,
			unitValue:     1,
			quantity:      1,
			ExpectedError: entity.InventoryProviderError,
		},
		{
			TestName:      "invalid unit value",
			product:       product,
			provider:      provider,
			storage:       storage,
			unitValue:     -1,
			quantity:      1,
			ExpectedError: entity.InventoryUnitValueError,
		},
		{
			TestName:      "invalid quantity",
			product:       product,
			provider:      provider,
			storage:       storage,
			unitValue:     1,
			quantity:      -1,
			ExpectedError: entity.InventoryQuantityError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			_, err := entity.NewInventory(tc.product, tc.storage, tc.provider, tc.unitValue, tc.quantity)
			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("expected error: %v\ngot: %v", tc.ExpectedError, err)
			}
		})
	}
}
