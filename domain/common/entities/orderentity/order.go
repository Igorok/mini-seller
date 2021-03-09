// Package orderentity - entity for order
package orderentity

import "time"

// Order - entity of order
type Order struct {
	ID               string       `validate:"omitempty"`
	IDCustomer       string       `validate:"omitempty"`
	IDOrganization   string       `validate:"required"`
	Basket           []BasketItem `validate:"required,gt=0"`
	Type             string       `validate:"required"`
	Status           string       `validate:"required"`
	Price            float64      `validate:"required,gt=0"`
	DateCreate       time.Time    `validate:"required"`
	Comment          string       `validate:"omitempty"`
	DateCustomerWish time.Time    `validate:"required"`
	Address          string       `validate:"required"`
}

// BasketItem - entity of product in basket
type BasketItem struct {
	ID    string  `validate:"required"`
	Price float64 `validate:"required,gt=0"`
	Count int     `validate:"required,gt=0"`
}
