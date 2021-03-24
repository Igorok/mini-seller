package schemas

import (
	"mini-seller/domain/packages/customer/catalogpkg"
	"mini-seller/domain/packages/customer/productpkg"
	"mini-seller/web/customer/controllers/catalogcontroller"
	"mini-seller/web/customer/controllers/productcontroller"

	"github.com/graphql-go/graphql"
)

// GetSchema - return graphql schema
func GetSchema(
	productUseCase productpkg.IUseCase,
	catalogUseCase catalogpkg.IUseCase,
) (graphql.Schema, error) {
	// initialization of controllers
	productController := productcontroller.NewProductController(productUseCase)
	// catalogController := catalogcontroller.NewCatalogController(catalogUseCase)

	organizationListResolver := catalogcontroller.NewOrganizationListResolver(catalogUseCase)

	fields := graphql.Fields{
		"getProductList":   productController.GetProductList(),
		"getProductDetail": productController.GetProductDetail(),

		"getOrganizationList": organizationListResolver.GetOrganizationList(),

		// "catalogGetOrganizationList":   catalogController.GetOrganizationList(),
		// "catalogGetOrganizationDetail": catalogController.GetOrganizationDetail(),
		// "catalogGetCategoryList":       catalogController.GetCategoryList(),
		// "catalogGetCategoryDetail":     catalogController.GetCategoryDetail(),
		// "catalogGetProductList":        catalogController.GetProductList(),
		// "catalogGetProductDetail":      catalogController.GetProductDetail(),
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
