package entity

import (
	"errors"
	"github.com/google/uuid"
	"inventory_api/entity/enum"
	"time"
)

var (
	ProductNameError        = errors.New("invalid name to product")
	ProductDescriptionError = errors.New("invalid description to product")
	ProductCategoryError    = errors.New("invalid category to product")
)

type Product struct {
	base
	name        string
	description string
	category    enum.Category
}

func NewProduct(name, description, category string) (Product, error) {
	//todo: Adicionar mais regras de negócio.
	if name == "" {
		return Product{}, ProductNameError
	}
	if description == "" {
		return Product{}, ProductDescriptionError
	}
	if !enum.ValidateCategory(category) {
		return Product{}, ProductCategoryError
	}

	return Product{
		base: base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		name:        name,
		description: description,
		category:    enum.Category(category),
	}, nil
}
