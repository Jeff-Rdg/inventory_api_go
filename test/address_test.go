package test

import (
	"errors"
	"inventory_api/entity"
	"testing"
)

func TestAddress_NewAddress(t *testing.T) {
	type TestCase struct {
		TestName      string
		city          string
		state         string
		country       string
		number        string
		road          string
		neighborhood  string
		ExpectedError error
	}

	testCases := []TestCase{
		{
			TestName:      "Invalid city",
			city:          "",
			state:         "Test",
			country:       "Test",
			number:        "Test",
			road:          "Test",
			neighborhood:  "Test",
			ExpectedError: entity.AddressError,
		},
		{
			TestName:      "Invalid state",
			city:          "Test",
			state:         "",
			country:       "Test",
			number:        "Test",
			road:          "Test",
			neighborhood:  "Test",
			ExpectedError: entity.AddressError,
		},
		{
			TestName:      "Valid address",
			city:          "Test",
			state:         "Test",
			country:       "Test",
			number:        "Test",
			road:          "Test",
			neighborhood:  "Test",
			ExpectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			_, err := entity.NewAddress(tc.city, tc.state, tc.country, tc.number, tc.road, tc.neighborhood)
			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("expected error: %v\ngot: %v", tc.ExpectedError, err)
			}
		})
	}
}
