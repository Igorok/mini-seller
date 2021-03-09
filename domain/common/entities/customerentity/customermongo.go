// Package customer - customer
package customerentity

import "go.mongodb.org/mongo-driver/bson/primitive"

// CustomerMongo - model for database
type CustomerMongo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string
	Sname  string
	Email  string
	Phone  string
	Status string
}

// ToCustomer - convert model to entity
func ToCustomer(cm *CustomerMongo) *Customer {
	ID := ""
	if !cm.ID.IsZero() {
		ID = cm.ID.Hex()
	}

	return &Customer{
		ID:     ID,
		Name:   cm.Name,
		Sname:  cm.Sname,
		Email:  cm.Email,
		Phone:  cm.Phone,
		Status: cm.Status,
	}
}

// ToCustomerMongo - convert model to entity
func ToCustomerMongo(c *Customer) *CustomerMongo {
	var ID primitive.ObjectID
	if c.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(c.ID)
	}

	return &CustomerMongo{
		ID:     ID,
		Name:   c.Name,
		Sname:  c.Sname,
		Email:  c.Email,
		Phone:  c.Phone,
		Status: c.Status,
	}
}
