package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/catalogentity"
	"mini-seller/domain/packages/catalogpkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (cRepo Repository) GetProductList(contx context.Context, id_organization string, id_category string) ([]*catalogentity.ProductInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// query
	query := bson.M{"status": catalogpkg.StatusActive}

	if id_organization != "" {
		IDOrganization, err := primitive.ObjectIDFromHex(id_organization)
		if err != nil {
			return nil, err
		}
		query["_id_org"] = IDOrganization
	}
	if id_category != "" {
		IDCategory, err := primitive.ObjectIDFromHex(id_category)
		if err != nil {
			return nil, err
		}
		query["_id_cat"] = IDCategory
	}

	// sort
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	// request
	cursor, err := cRepo.db.Collection("products").Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}

	// format data from cursor
	products := make([]*catalogentity.ProductInfo, 0)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		prodMongo := catalogentity.ProductInfoMongo{}
		err := cursor.Decode(&prodMongo)
		if err != nil {
			return nil, err
		}
		product := catalogentity.ToProductInfo(prodMongo)
		products = append(products, &product)
	}

	// answer
	return products, nil
}

func (cRepo Repository) GetProductDetail(contx context.Context, id string) (*catalogentity.ProductInfo, error) {
	// convert id to bson
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// context
	prodMongo := catalogentity.ProductInfoMongo{}
	err = cRepo.db.Collection("products").FindOne(ctx, bson.M{"_id": ID, "status": catalogpkg.StatusActive}).Decode(&prodMongo)
	if err != nil {
		return nil, err
	}

	// convert to entity
	product := catalogentity.ToProductInfo(prodMongo)

	// answer
	return &product, nil
}
