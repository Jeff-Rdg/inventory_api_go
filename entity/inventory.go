package entity

type Inventory struct {
	Base
	Product
	Storage
	Provider
	unitValue float64
	quantity  int
}
