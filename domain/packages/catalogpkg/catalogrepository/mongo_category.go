package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/packages/catalogpkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (cRepo Repository) GetCategoryList(contx context.Context, ids []string) ([]*productcategoryentity.ProductCategory, error) {
	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// query
	query := bson.M{"status": catalogpkg.StatusActive}
	if ids != nil && len(ids) > 0 {
		ids_cat := make([]primitive.ObjectID, len(ids))
		for i, id := range ids {
			id_cat, err_id := primitive.ObjectIDFromHex(id)
			if err_id != nil {
				return nil, err_id
			}
			ids_cat[i] = id_cat
		}
		query["_id"] = bson.M{"$in": ids_cat}
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})

	// request
	cursor, err := cRepo.db.Collection("product_categories").Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}

	// format data from cursor
	categories := make([]*productcategoryentity.ProductCategory, 0)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		catMongo := productcategoryentity.ProductCategoryMongo{}
		err := cursor.Decode(&catMongo)
		if err != nil {
			return nil, err
		}
		category := productcategoryentity.ToEntity(&catMongo)
		categories = append(categories, category)
	}

	// answer
	return categories, nil
}

func (cRepo Repository) GetCategoryDetail(contx context.Context, id string) (*productcategoryentity.ProductCategory, error) {
	// convert id to bson
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// context
	catMongo := productcategoryentity.ProductCategoryMongo{}
	err = cRepo.db.Collection("product_categories").FindOne(ctx, bson.M{"_id": ID, "status": catalogpkg.StatusActive}).Decode(&catMongo)
	if err != nil {
		return nil, err
	}

	// convert to entity
	category := productcategoryentity.ToEntity(&catMongo)

	// answer
	return category, nil
}
