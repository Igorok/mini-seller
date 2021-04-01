// Package organizationentity - entities for organization
package organizationentity

import "go.mongodb.org/mongo-driver/bson/primitive"

// OrganizationMongo - model for mongo database
type OrganizationMongo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Email       string
	Phone       string
	Status      string
	IDsCategory []primitive.ObjectID `bson:"_ids_cat,omitempty"`
}

// ToEntity - method to convert model to entity
func ToEntity(om *OrganizationMongo) *Organization {
	ID := ""
	if !om.ID.IsZero() {
		ID = om.ID.Hex()
	}
	IDsCategory := make([]string, 0)
	if om.IDsCategory != nil && len(om.IDsCategory) > 0 {
		for _, id := range om.IDsCategory {
			IDsCategory = append(IDsCategory, id.Hex())
		}
	}

	return &Organization{
		ID:          ID,
		Name:        om.Name,
		Email:       om.Email,
		Phone:       om.Phone,
		Status:      om.Status,
		IDsCategory: IDsCategory,
	}
}

// ToMongo - method to convert entity to model
func ToMongo(o *Organization) *OrganizationMongo {
	var ID primitive.ObjectID
	if o.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(o.ID)
	}

	IDsCategory := make([]primitive.ObjectID, 0)
	if o.IDsCategory != nil && len(o.IDsCategory) > 0 {
		for _, id := range o.IDsCategory {
			if id != "" {
				idMongo, _ := primitive.ObjectIDFromHex(id)
				IDsCategory = append(IDsCategory, idMongo)
			}
		}
	}

	return &OrganizationMongo{
		ID:          ID,
		Name:        o.Name,
		Email:       o.Email,
		Phone:       o.Phone,
		Status:      o.Status,
		IDsCategory: IDsCategory,
	}
}
