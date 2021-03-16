package schemas

import (
	"mini-seller/domain/packages/customer/catalog"
	"mini-seller/web/customer/controllers/catalogcontroller"

	"github.com/graphql-go/graphql"
)

// GetSchema - return graphql schema
func GetSchema(
	catalogUseCase catalog.IUseCase,
) (graphql.Schema, error) {
	// initialization of controllers
	catalogController := catalogcontroller.NewCatalogController(catalogUseCase)

	fields := graphql.Fields{
		"getCatalog": catalogController.GetCatalog(),
	}
	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	return graphql.NewSchema(schemaConfig)
}
