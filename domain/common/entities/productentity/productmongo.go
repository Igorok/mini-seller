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

// ProductForCatalogMongo - model for mongo database
type ProductForCatalogMongo struct {
	ID           primitive.ObjectID `bson:"_id"`
	Category     CategoryForCatalogMongo
	Organization OrganizationForCatalogMongo
	Name         string
	Price        int
	Count        int
	Status       string
}

// CategoryForCatalogMongo - model for mongo database
type CategoryForCatalogMongo struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string
}

// OrganizationForCatalogMongo - model for mongo database
type OrganizationForCatalogMongo struct {
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

// ToCategoryForCatalog - convert model to entity
func ToCategoryForCatalog(catForCatM *CategoryForCatalogMongo) *CategoryForCatalog {
	return &CategoryForCatalog{
		ID:   catForCatM.ID.Hex(),
		Name: catForCatM.Name,
	}
}

// ToCategoryForCatalogMongo - convert model to entity
func ToCategoryForCatalogMongo(catForCat *CategoryForCatalog) *CategoryForCatalogMongo {
	ID, _ := primitive.ObjectIDFromHex(catForCat.ID)
	return &CategoryForCatalogMongo{
		ID:   ID,
		Name: catForCat.Name,
	}
}

// ToOrganizationForCatalog - convert model to entity
func ToOrganizationForCatalog(orgForCatM *OrganizationForCatalogMongo) *OrganizationForCatalog {
	return &OrganizationForCatalog{
		ID:    orgForCatM.ID.Hex(),
		Name:  orgForCatM.Name,
		Phone: orgForCatM.Phone,
		Email: orgForCatM.Email,
	}
}

// ToOrganizationForCatalogMongo - convert model to entity
func ToOrganizationForCatalogMongo(orgForCat *OrganizationForCatalog) *OrganizationForCatalogMongo {
	ID, _ := primitive.ObjectIDFromHex(orgForCat.ID)
	return &OrganizationForCatalogMongo{
		ID:    ID,
		Name:  orgForCat.Name,
		Phone: orgForCat.Phone,
		Email: orgForCat.Email,
	}
}

// ToProductForCatalog - convert model to entity
func ToProductForCatalog(prodForCatM *ProductForCatalogMongo) *ProductForCatalog {
	Category := ToCategoryForCatalog(&prodForCatM.Category)
	Organization := ToOrganizationForCatalog(&prodForCatM.Organization)
	return &ProductForCatalog{
		ID:           prodForCatM.ID.Hex(),
		Name:         prodForCatM.Name,
		Price:        prodForCatM.Price,
		Count:        prodForCatM.Count,
		Status:       prodForCatM.Status,
		Category:     *Category,
		Organization: *Organization,
	}
}

// ToProductForCatalogMongo - convert model to entity
func ToProductForCatalogMongo(prodForCat *ProductForCatalog) *ProductForCatalogMongo {
	ID, _ := primitive.ObjectIDFromHex(prodForCat.ID)
	Category := ToCategoryForCatalogMongo(&prodForCat.Category)
	Organization := ToOrganizationForCatalogMongo(&prodForCat.Organization)
	return &ProductForCatalogMongo{
		ID:           ID,
		Name:         prodForCat.Name,
		Price:        prodForCat.Price,
		Count:        prodForCat.Count,
		Status:       prodForCat.Status,
		Category:     *Category,
		Organization: *Organization,
	}
}
