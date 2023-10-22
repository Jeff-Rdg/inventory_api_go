package entity

import "inventory_api/entity/enum"

type Product struct {
	Base
	name        string
	description string
	category    enum.Category
}
