// Package orderevententity - entities for events from order
package orderevententity

import "time"

// OrderEvent - events from order
type OrderEvent struct {
	ID         string    `validate:"omitempty"`
	IDOrder    string    `validate:"required"`
	IDEmployee string    `validate:"required"`
	Type       string    `validate:"required"`
	Status     string    `validate:"required"`
	Date       time.Time `validate:"required"`
}
