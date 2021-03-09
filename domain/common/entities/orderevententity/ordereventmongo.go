// Package orderevententity - entities for events from order
package orderevententity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderEventMongo - model for database
type OrderEventMongo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	IDOrder    primitive.ObjectID `bson:"_id_order"`
	IDEmployee primitive.ObjectID `bson:"_id_empl"`
	Type       string
	Status     string
	Date       time.Time
}

// ToEntity - convert model to entity
func ToEntity(oem *OrderEventMongo) *OrderEvent {
	ID := ""
	if !oem.ID.IsZero() {
		ID = oem.ID.Hex()
	}

	IDOrder := oem.IDOrder.Hex()
	IDEmployee := oem.IDEmployee.Hex()

	return &OrderEvent{
		ID:         ID,
		IDOrder:    IDOrder,
		IDEmployee: IDEmployee,
		Type:       oem.Type,
		Status:     oem.Status,
		Date:       oem.Date,
	}
}

// ToMongo - convert entity to model
func ToMongo(oe *OrderEvent) *OrderEventMongo {
	var ID primitive.ObjectID
	if oe.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(oe.ID)
	}

	IDOrder, _ := primitive.ObjectIDFromHex(oe.IDOrder)
	IDEmployee, _ := primitive.ObjectIDFromHex(oe.IDEmployee)

	return &OrderEventMongo{
		ID:         ID,
		IDOrder:    IDOrder,
		IDEmployee: IDEmployee,
		Type:       oe.Type,
		Status:     oe.Status,
		Date:       oe.Date,
	}
}
