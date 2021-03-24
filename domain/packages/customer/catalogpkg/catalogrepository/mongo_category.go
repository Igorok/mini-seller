package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/catalogentity"
	"mini-seller/domain/packages/customer/catalogpkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (cRepo Repository) GetCategoryList(contx context.Context) ([]*catalogentity.CategoryInfo, error) {
	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// request
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	cursor, err := cRepo.db.Collection("product_categories").Find(ctx, bson.M{"status": catalogpkg.StatusActive}, findOptions)
	if err != nil {
		return nil, err
	}

	// format data from cursor
	categories := make([]*catalogentity.CategoryInfo, 0)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		catMongo := catalogentity.CategoryInfoMongo{}
		err := cursor.Decode(&catMongo)
		if err != nil {
			return nil, err
		}
		category := catalogentity.ToCategoryInfo(catMongo)
		categories = append(categories, &category)
	}

	// answer
	return categories, nil
}

func (cRepo Repository) GetCategoryDetail(contx context.Context, id string) (*catalogentity.CategoryInfo, error) {
	// convert id to bson
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// context
	catMongo := catalogentity.CategoryInfoMongo{}
	err = cRepo.db.Collection("product_categories").FindOne(ctx, bson.M{"_id": ID, "status": catalogpkg.StatusActive}).Decode(&catMongo)
	if err != nil {
		return nil, err
	}

	// convert to entity
	category := catalogentity.ToCategoryInfo(catMongo)

	// answer
	return &category, nil
}
