package test

import (
	"errors"
	"inventory_api/entity"
	"testing"
)

func TestStorage_NewStorage(t *testing.T) {
	type TestCase struct {
		TestName      string
		Location      string
		Description   string
		ExpectedError error
	}

	testCases := []TestCase{
		{
			TestName:      "Invalid Location",
			Location:      "",
			Description:   "Test",
			ExpectedError: entity.StorageLocationError,
		},
		{
			TestName:      "Valid Location",
			Location:      "Prateleira",
			Description:   "",
			ExpectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			_, err := entity.NewStorage(tc.Location, tc.Description)
			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("expected error: %v\ngot: %v", tc.ExpectedError, err)
			}
		})
	}
}
