package test

import (
	"errors"
	"inventory_api/entity"
	"testing"
)

func TestProvider_NewProvider(t *testing.T) {
	address, _ := entity.NewAddress("Test", "Test", "Test", "Test", "Test", "Test")

	type TestCase struct {
		TestName      string
		address       entity.Address
		corporateName string
		cnpj          string
		ExpectedError error
	}

	testCases := []TestCase{
		{
			TestName:      "Invalid address",
			address:       entity.Address{},
			corporateName: "Test",
			cnpj:          "Test",
			ExpectedError: entity.ProviderAddressError,
		},
		{
			TestName:      "Valid provider",
			address:       address,
			corporateName: "Test",
			cnpj:          "Test",
			ExpectedError: nil,
		},
		{
			TestName:      "Invalid corporateName",
			address:       address,
			corporateName: "",
			cnpj:          "Test",
			ExpectedError: entity.ProviderCorporateNameError,
		},
		{
			TestName:      "Invalid cnpj",
			address:       address,
			corporateName: "Test",
			cnpj:          "",
			ExpectedError: entity.ProviderCnpjError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			_, err := entity.NewProvider(tc.address, tc.corporateName, tc.cnpj)
			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("expected error: %v\ngot: %v", tc.ExpectedError, err)
			}
		})
	}
}
