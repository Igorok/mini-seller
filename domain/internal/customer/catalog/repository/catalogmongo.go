package catalog

import (
	"context"
	"log"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/internal/customer/catalog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// commonPipeline - for detail and list info
var commonPipeline = []bson.D{
	{{"$sort", bson.M{
		"_id": 1,
	}}},

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

// commonAggregate - request for aggregation of products info
func commonAggregate(db *mongo.Database, pipeline []bson.D) ([]*productentity.ProductForCatalog, error) {
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

// GetCatalog - get list of products from mongodb
func (cRepo Repository) GetCatalog(ctx context.Context, skip, limit int) ([]*productentity.ProductForCatalog, error) {
	// merge query
	pipeline := []bson.D{
		{{"$match", bson.M{
			"status": catalog.StatusActive,
		}}},
	}
	pipeline = append(pipeline, commonPipeline...)

	// get data
	return commonAggregate(cRepo.db, pipeline)
}

// GetProductDetail - get details for product from mongodb
func (cRepo Repository) GetProductDetail(ctx context.Context, id string) (*productentity.ProductForCatalog, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// merge query
	pipeline := []bson.D{
		{{"$match", bson.M{
			"_id": ID,
		}}},
	}
	pipeline = append(pipeline, commonPipeline...)

	// get data
	products, err := commonAggregate(cRepo.db, pipeline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return products[0], nil
}
