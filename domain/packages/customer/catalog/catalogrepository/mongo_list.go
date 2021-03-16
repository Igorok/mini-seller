package catalogrepository

import (
	"context"
	"log"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/customer/catalog"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func catalogList(db *mongo.Database, skip, limit int) ([]*productentity.ProductForCatalog, error) {
	// build query
	pipeline := []bson.D{
		{{"$match", bson.M{
			"status": catalog.StatusActive,
		}}},
		sortPipeline,
		{{"$skip", skip}},
		{{"$limit", limit}},
	}
	pipeline = append(pipeline, lookupPipeline...)

	// get data
	return aggregateProducts(db, pipeline)
}

func catalogCount(db *mongo.Database) (int64, error) {
	// request
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	countQ := bson.M{"status": catalog.StatusActive}
	return db.Collection("products").CountDocuments(ctx, countQ)
}

// GetCatalog - get list of products from mongodb
func (cRepo Repository) GetCatalog(ctx context.Context, skip, limit int) ([]*productentity.ProductForCatalog, int64, error) {
	var productList []*productentity.ProductForCatalog
	var count int64
	var errlist, errcount error
	var wg sync.WaitGroup

	// get products
	wg.Add(1)
	go func() {
		defer wg.Done()
		productList, errlist = catalogList(cRepo.db, skip, limit)
	}()

	// get count
	wg.Add(1)
	go func() {
		defer wg.Done()
		count, errcount = catalogCount(cRepo.db)
	}()

	// wait goroutines
	wg.Wait()

	if errlist != nil {
		log.Fatal(errlist)
		return nil, 0, errlist
	}
	if errcount != nil {
		log.Fatal(errcount)
		return nil, 0, errcount
	}

	return productList, count, nil
}
