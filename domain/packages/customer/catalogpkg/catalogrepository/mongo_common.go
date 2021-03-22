package catalogrepository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	db *mongo.Database
}

// NewCatalogRepository - constructor for catalog repository
func NewCatalogRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}
