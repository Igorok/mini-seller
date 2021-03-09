// Package organizationentity - entities for organization
package organizationentity

import "go.mongodb.org/mongo-driver/bson/primitive"

// OrganizationMongo - model for mongo database
type OrganizationMongo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string
	Email  string
	Phone  string
	Status string
}

// ToEntity - method to convert model to entity
func ToEntity(om *OrganizationMongo) *Organization {
	ID := ""
	if !om.ID.IsZero() {
		ID = om.ID.Hex()
	}

	return &Organization{
		ID:     ID,
		Name:   om.Name,
		Email:  om.Email,
		Phone:  om.Phone,
		Status: om.Status,
	}
}

// ToMongo - method to convert entity to model
func ToMongo(o *Organization) *OrganizationMongo {
	var ID primitive.ObjectID
	if o.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(o.ID)
	}

	return &OrganizationMongo{
		ID:     ID,
		Name:   o.Name,
		Email:  o.Email,
		Phone:  o.Phone,
		Status: o.Status,
	}
}
