package test

import (
	"errors"
	"inventory_api/entity"
	"testing"
)

func TestProduct_NewProduct(t *testing.T) {

	type TestCase struct {
		TestName      string
		name          string
		description   string
		category      string
		ExpectedError error
	}

	testCases := []TestCase{
		{
			TestName:      "Invalid name",
			name:          "",
			description:   "Test",
			category:      "OTHER",
			ExpectedError: entity.ProductNameError,
		},
		{
			TestName:      "Valid product",
			name:          "test",
			description:   "Test",
			category:      "OTHER",
			ExpectedError: nil,
		},
		{
			TestName:      "Invalid description",
			name:          "test",
			description:   "",
			category:      "OTHER",
			ExpectedError: entity.ProductDescriptionError,
		},
		{
			TestName:      "Invalid category",
			name:          "test",
			description:   "test",
			category:      "OTHERs",
			ExpectedError: entity.ProductCategoryError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			_, err := entity.NewProduct(tc.name, tc.description, tc.category)
			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("expected error: %v\ngot: %v", tc.ExpectedError, err)
			}
		})
	}
}
