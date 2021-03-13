package repository

import (
	"context"
	"log"
	"mini-seller/domain/common/entities/productentity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		sortPipeline,
	}
	pipeline = append(pipeline, lookupPipeline...)

	// get data
	products, err := aggregateProducts(cRepo.db, pipeline)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return products[0], nil
}
