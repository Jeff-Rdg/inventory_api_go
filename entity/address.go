package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	AddressError = errors.New("invalid Address")
)

type Address struct {
	Base
	city         string
	state        string
	country      string
	number       string
	road         string
	neighborhood string
}

func NewAddress(city, state, country, number, road, neighborhood string) (Address, error) {
	if city == "" || state == "" || country == "" || number == "" || road == "" || neighborhood == "" {
		return Address{}, AddressError
	}

	return Address{
		Base: Base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		city:         city,
		state:        state,
		country:      country,
		number:       number,
		road:         road,
		neighborhood: neighborhood,
	}, nil
}
