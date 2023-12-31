package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	StorageLocationError = errors.New("invalid Location")
)

type Storage struct {
	base
	location    string
	description string
}

func NewStorage(location, description string) (Storage, error) {
	if location == "" {
		return Storage{}, StorageLocationError
	}

	return Storage{
		base: base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		location:    location,
		description: description,
	}, nil
}
