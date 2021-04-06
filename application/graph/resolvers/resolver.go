package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"mini-seller/application/graph/model"
	"mini-seller/domain/packages/catalogpkg"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CatalogUseCase catalogpkg.IUseCase

	todos         []*model.Todo
	organizations []*model.Organization
}
