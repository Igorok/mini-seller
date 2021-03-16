package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository - managing of catalog data
type Repository struct {
	db *mongo.Database
}

// NewCatalogRepository - constructor for catalog repository
func NewCatalogRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}

// sortPipeline - sort by id
var sortPipeline = bson.D{{"$sort", bson.M{
	"_id": 1,
}}}

// lookupPipeline - join organizations and product_categories
var lookupPipeline = []bson.D{
	{{"$lookup", bson.M{
		"from":         "organizations",
		"localField":   "_id_org",
		"foreignField": "_id",
		"as":           "Organization",
	}}},
	{{"$lookup", bson.M{
		"from":         "product_categories",
		"localField":   "_id_cat",
		"foreignField": "_id",
		"as":           "Category",
	}}},

	{{"$unwind", "$Organization"}},
	{{"$unwind", "$Category"}},
}

// aggregateProducts - request for aggregation products info
func aggregateProducts(db *mongo.Database, pipeline []bson.D) ([]*productentity.ProductForCatalog, error) {
	// request
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.Collection("products").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// format data from cursor
	products := make([]*productentity.ProductForCatalog, 0)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		prodMongo := new(productentity.ProductForCatalogMongo)
		err := cursor.Decode(&prodMongo)
		if err != nil {
			return nil, err
		}
		prod := productentity.ToProductForCatalog(prodMongo)
		products = append(products, prod)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
