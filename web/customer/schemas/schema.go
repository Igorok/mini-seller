package schemas

import (
	"mini-seller/domain/packages/customer/productpkg"
	"mini-seller/web/customer/controllers/productcontroller"

	"github.com/graphql-go/graphql"
)

// GetSchema - return graphql schema
func GetSchema(
	productUseCase productpkg.IUseCase,
) (graphql.Schema, error) {
	// initialization of controllers
	productController := productcontroller.NewProductController(productUseCase)

	fields := graphql.Fields{
		"getProductList": productController.GetProductList(),
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
