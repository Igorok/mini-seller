// Package productentity - entities for catalog package
package productentity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductMongo - model for mongo database
type ProductMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	IDCategory     primitive.ObjectID `bson:"_id_cat"`
	IDOrganization primitive.ObjectID `bson:"_id_org"`
	Name           string
	Price          int
	Count          int
	Status         string
}

// ProductForListMongo - model for mongo database
type ProductForListMongo struct {
	ID           primitive.ObjectID `bson:"_id"`
	Category     CategoryForProductMongo
	Organization OrganizationForProductMongo
	Name         string
	Price        int
	Count        int
	Status       string
}

// CategoryForProductMongo - model for mongo database
type CategoryForProductMongo struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string
}

// OrganizationForProductMongo - model for mongo database
type OrganizationForProductMongo struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string
	Phone string
	Email string
}

// ToProduct - method to convert model to entity
func ToProduct(pm *ProductMongo) *Product {
	ID := ""
	if !pm.ID.IsZero() {
		ID = pm.ID.Hex()
	}

	return &Product{
		ID:             ID,
		IDCategory:     pm.IDCategory.Hex(),
		IDOrganization: pm.IDOrganization.Hex(),
		Name:           pm.Name,
		Price:          pm.Price,
		Count:          pm.Count,
		Status:         pm.Status,
	}
}

// ToProductMongo - method to convert entity to model
func ToProductMongo(p *Product) *ProductMongo {
	var ID primitive.ObjectID
	if p.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(p.ID)
	}
	idCategory, _ := primitive.ObjectIDFromHex(p.IDCategory)
	idOrganization, _ := primitive.ObjectIDFromHex(p.IDOrganization)

	return &ProductMongo{
		ID:             ID,
		IDCategory:     idCategory,
		IDOrganization: idOrganization,
		Name:           p.Name,
		Price:          p.Price,
		Count:          p.Count,
		Status:         p.Status,
	}
}

// ToCategoryForProduct - convert model to entity
func ToCategoryForProduct(catForCatM *CategoryForProductMongo) *CategoryForProduct {
	return &CategoryForProduct{
		ID:   catForCatM.ID.Hex(),
		Name: catForCatM.Name,
	}
}

// ToCategoryForProductMongo - convert model to entity
func ToCategoryForProductMongo(catForCat *CategoryForProduct) *CategoryForProductMongo {
	ID, _ := primitive.ObjectIDFromHex(catForCat.ID)
	return &CategoryForProductMongo{
		ID:   ID,
		Name: catForCat.Name,
	}
}

// ToOrganizationForProduct - convert model to entity
func ToOrganizationForProduct(orgForCatM *OrganizationForProductMongo) *OrganizationForProduct {
	return &OrganizationForProduct{
		ID:    orgForCatM.ID.Hex(),
		Name:  orgForCatM.Name,
		Phone: orgForCatM.Phone,
		Email: orgForCatM.Email,
	}
}

// ToOrganizationForProductMongo - convert model to entity
func ToOrganizationForProductMongo(orgForCat *OrganizationForProduct) *OrganizationForProductMongo {
	ID, _ := primitive.ObjectIDFromHex(orgForCat.ID)
	return &OrganizationForProductMongo{
		ID:    ID,
		Name:  orgForCat.Name,
		Phone: orgForCat.Phone,
		Email: orgForCat.Email,
	}
}

// ToProductForList - convert model to entity
func ToProductForList(prodForCatM *ProductForListMongo) *ProductForList {
	Category := ToCategoryForProduct(&prodForCatM.Category)
	Organization := ToOrganizationForProduct(&prodForCatM.Organization)
	return &ProductForList{
		ID:           prodForCatM.ID.Hex(),
		Name:         prodForCatM.Name,
		Price:        prodForCatM.Price,
		Count:        prodForCatM.Count,
		Status:       prodForCatM.Status,
		Category:     *Category,
		Organization: *Organization,
	}
}

// ToProductForListMongo - convert model to entity
func ToProductForListMongo(prodForCat *ProductForList) *ProductForListMongo {
	ID, _ := primitive.ObjectIDFromHex(prodForCat.ID)
	Category := ToCategoryForProductMongo(&prodForCat.Category)
	Organization := ToOrganizationForProductMongo(&prodForCat.Organization)
	return &ProductForListMongo{
		ID:           ID,
		Name:         prodForCat.Name,
		Price:        prodForCat.Price,
		Count:        prodForCat.Count,
		Status:       prodForCat.Status,
		Category:     *Category,
		Organization: *Organization,
	}
}
