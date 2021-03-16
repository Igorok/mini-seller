package catalogcontroller

import (
	"mini-seller/domain/packages/customer/catalog"
	"mini-seller/web/customer/fields/catalogfield"

	"github.com/graphql-go/graphql"
)

// CatalogController - web controller
type CatalogController struct {
	catalogUseCase catalog.IUseCase
}

// NewCatalogController - constructor
func NewCatalogController(catalogUseCase catalog.IUseCase) *CatalogController {
	return &CatalogController{catalogUseCase: catalogUseCase}
}

// Products - products list
func (controller CatalogController) GetCatalog() *graphql.Field {
	return &graphql.Field{
		Type: catalogfield.Catalog,
		Args: graphql.FieldConfigArgument{
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			skip, ok := p.Args["skip"].(int)
			if !ok {
				return nil, catalog.ErrCatalogLimit
			}
			limit, ok := p.Args["limit"].(int)
			if !ok {
				return nil, catalog.ErrCatalogLimit
			}

			return controller.catalogUseCase.GetCatalog(p.Context, skip, limit)
		},
	}
}
