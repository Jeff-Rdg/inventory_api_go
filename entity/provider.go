package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ProviderAddressError       = errors.New("invalid address to provider")
	ProviderCorporateNameError = errors.New("invalid corporate Name to provider")
	ProviderCnpjError          = errors.New("invalid cnpj to provider")
)

type Provider struct {
	base
	address       Address
	corporateName string
	cnpj          string
}

func NewProvider(address Address, corporateName, cnpj string) (Provider, error) {
	if (address == Address{}) {
		return Provider{}, ProviderAddressError
	}
	if corporateName == "" {
		return Provider{}, ProviderCorporateNameError
	}
	if cnpj == "" {
		return Provider{}, ProviderCnpjError
	}

	return Provider{
		base: base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		address:       address,
		corporateName: corporateName,
		cnpj:          cnpj,
	}, nil
}
