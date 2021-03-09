// Package customerentity - customer
package customerentity

// Customer - entity for customer
type Customer struct {
	ID     string `validate:"omitempty"`
	Name   string `validate:"required"`
	Sname  string `validate:"omitempty"`
	Email  string `validate:"required,email"`
	Phone  string `validate:"required"`
	Status string `validate:"required"`
}
