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

// GetOrganizationList - list with active organizations
func (cRepo Repository) GetOrganizationList(contx context.Context) ([]*catalogentity.OrganizationInfo, error) {
	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// request
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	cursor, err := cRepo.db.Collection("organizations").Find(ctx, bson.M{"status": catalogpkg.StatusActive}, findOptions)
	if err != nil {
		return nil, err
	}

	// format data from cursor
	organizations := make([]*catalogentity.OrganizationInfo, 0)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		orgMongo := catalogentity.OrganizationInfoMongo{}
		err := cursor.Decode(&orgMongo)
		if err != nil {
			return nil, err
		}
		org := catalogentity.ToOrganizationInfo(orgMongo)
		organizations = append(organizations, &org)
	}

	// answer
	return organizations, nil
}

/*
GetOrganizationDetail - get detail info about organization
@param {String} id - id of organization
@return catalogentity.OrganizationInfo, error - detail info and error
*/
func (cRepo Repository) GetOrganizationDetail(contx context.Context, id string) (*catalogentity.OrganizationInfo, error) {
	// convert id to bson
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// context
	orgMongo := catalogentity.OrganizationInfoMongo{}
	err = cRepo.db.Collection("organizations").FindOne(ctx, bson.M{"_id": ID, "status": catalogpkg.StatusActive}).Decode(&orgMongo)
	if err != nil {
		return nil, err
	}

	// convert to entity
	organization := catalogentity.ToOrganizationInfo(orgMongo)

	// answer
	return &organization, nil
}
