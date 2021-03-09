// Package productcategoryentity - entities for catalog package
package productcategoryentity

import "go.mongodb.org/mongo-driver/bson/primitive"

// ProductCategoryMongo - model for mongo database
type ProductCategoryMongo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string
	Status string
}

// ToEntity - method to convert model to entity
func ToEntity(pm *ProductCategoryMongo) *ProductCategory {
	ID := ""
	if !pm.ID.IsZero() {
		ID = pm.ID.Hex()
	}
	return &ProductCategory{
		ID:     ID,
		Name:   pm.Name,
		Status: pm.Status,
	}
}

// ToMongo - method to convert entity to model
func ToMongo(p *ProductCategory) *ProductCategoryMongo {
	var ID primitive.ObjectID
	if p.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(p.ID)
	}

	return &ProductCategoryMongo{
		ID:     ID,
		Name:   p.Name,
		Status: p.Status,
	}
}
