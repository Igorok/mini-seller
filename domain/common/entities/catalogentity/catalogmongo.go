// Package catalogentity - entities for catalog
package catalogentity

import "go.mongodb.org/mongo-driver/bson/primitive"

// ProductInfoMongo - model for product
type ProductInfoMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	IDCategory     primitive.ObjectID `bson:"_id_cat"`
	IDOrganization primitive.ObjectID `bson:"_id_org"`
	Name           string
	Price          int
	Count          int
	Status         string
}

// CategoryInfoMongo - model for category of products
type CategoryInfoMongo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string
	Status string
}

// OrganizationInfoMongo - model for organization
type OrganizationInfoMongo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string
	Email  string
	Phone  string
	Status string
}

// ToProductInfo - convert model to entity
func ToProductInfo(pim ProductInfoMongo) ProductInfo {
	return ProductInfo{
		ID:             pim.ID.Hex(),
		IDCategory:     pim.ID.Hex(),
		IDOrganization: pim.ID.Hex(),
		Name:           pim.Name,
		Price:          pim.Price,
		Count:          pim.Count,
		Status:         pim.Status,
	}
}

// ToProductInfoMongo - convert entity to model
func ToProductInfoMongo(pi ProductInfo) ProductInfoMongo {
	id, _ := primitive.ObjectIDFromHex(pi.ID)
	idCategory, _ := primitive.ObjectIDFromHex(pi.IDCategory)
	idOrganization, _ := primitive.ObjectIDFromHex(pi.IDOrganization)

	return ProductInfoMongo{
		ID:             id,
		IDCategory:     idCategory,
		IDOrganization: idOrganization,
		Name:           pi.Name,
		Price:          pi.Price,
		Count:          pi.Count,
		Status:         pi.Status,
	}
}

// ToCategoryInfo - convert model to entity
func ToCategoryInfo(cim CategoryInfoMongo) CategoryInfo {
	return CategoryInfo{
		ID:     cim.ID.Hex(),
		Name:   cim.Name,
		Status: cim.Status,
	}
}

// ToCategoryInfoMongo - convert entity to model
func ToCategoryInfoMongo(ci CategoryInfo) CategoryInfoMongo {
	id, _ := primitive.ObjectIDFromHex(ci.ID)

	return CategoryInfoMongo{
		ID:     id,
		Name:   ci.Name,
		Status: ci.Status,
	}
}

// ToOrganizationInfo - convert model to entity
func ToOrganizationInfo(oim OrganizationInfoMongo) OrganizationInfo {
	return OrganizationInfo{
		ID:     oim.ID.Hex(),
		Name:   oim.Name,
		Email:  oim.Email,
		Phone:  oim.Phone,
		Status: oim.Status,
	}
}

// ToOrganizationInfoMongo - convert entity to model
func ToOrganizationInfoMongo(oi OrganizationInfo) OrganizationInfoMongo {
	id, _ := primitive.ObjectIDFromHex(oi.ID)

	return OrganizationInfoMongo{
		ID:     id,
		Name:   oi.Name,
		Email:  oi.Email,
		Phone:  oi.Phone,
		Status: oi.Status,
	}
}
