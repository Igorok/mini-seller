package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/catalogpkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (cRepo Repository) GetProductList(contx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// query
	query := bson.M{"status": catalogpkg.StatusActive}

	if ids_organization != nil && len(ids_organization) > 0 {
		ids_org := make([]primitive.ObjectID, len(ids_organization))
		for i, id := range ids_organization {
			id_org, err_id := primitive.ObjectIDFromHex(id)
			if err_id != nil {
				return nil, err_id
			}
			ids_org[i] = id_org
		}
		query["_id_org"] = bson.M{"$in": ids_org}
	}

	if ids_category != nil && len(ids_category) > 0 {
		ids_cat := make([]primitive.ObjectID, len(ids_category))
		for i, id := range ids_category {
			id_cat, err_id := primitive.ObjectIDFromHex(id)
			if err_id != nil {
				return nil, err_id
			}
			ids_cat[i] = id_cat
		}
		query["_id_cat"] = bson.M{"$in": ids_cat}
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
	products := make([]*productentity.Product, 0)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		prodMongo := productentity.ProductMongo{}
		err := cursor.Decode(&prodMongo)
		if err != nil {
			return nil, err
		}
		product := productentity.ToProduct(&prodMongo)
		products = append(products, product)
	}

	// answer
	return products, nil
}

func (cRepo Repository) GetProductDetail(contx context.Context, id string) (*productentity.Product, error) {
	// convert id to bson
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// context
	prodMongo := productentity.ProductMongo{}
	err = cRepo.db.Collection("products").FindOne(ctx, bson.M{"_id": ID, "status": catalogpkg.StatusActive}).Decode(&prodMongo)
	if err != nil {
		return nil, err
	}

	// convert to entity
	product := productentity.ToProduct(&prodMongo)

	// answer
	return product, nil
}
