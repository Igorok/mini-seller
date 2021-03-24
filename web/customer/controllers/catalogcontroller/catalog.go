package catalogcontroller

import (
	"mini-seller/domain/packages/customer/catalogpkg"
	"mini-seller/web/customer/fields/catalogfield"

	"github.com/graphql-go/graphql"
)

// CatalogController - web controller
type CatalogController struct {
	catalogUseCase catalogpkg.IUseCase
}

// NewCatalogController - constructor
func NewCatalogController(catalogUseCase catalogpkg.IUseCase) *CatalogController {
	controller := &CatalogController{catalogUseCase: catalogUseCase}

	return controller
}

/*
func (controller CatalogController) GetOrganizationList() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(catalogfield.CategoryInfo),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.catalogUseCase.GetOrganizationList(p.Context)
		},
	}
}
*/

func (controller CatalogController) GetOrganizationDetail() *graphql.Field {
	return &graphql.Field{
		Type: catalogfield.OrganizationInfo,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if !ok {
				return nil, catalogpkg.ErrOrganizationNotFound
			}
			return controller.catalogUseCase.GetOrganizationDetail(p.Context, id)
		},
	}
}

func (controller CatalogController) GetCategoryList() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(catalogfield.CategoryInfo),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.catalogUseCase.GetCategoryList(p.Context)
		},
	}
}

func (controller CatalogController) GetCategoryDetail() *graphql.Field {
	return &graphql.Field{
		Type: catalogfield.CategoryInfo,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if !ok {
				return nil, catalogpkg.ErrCategoryNotFound
			}
			return controller.catalogUseCase.GetCategoryDetail(p.Context, id)
		},
	}
}

func (controller CatalogController) GetProductList() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(catalogfield.ProductInfo),
		Args: graphql.FieldConfigArgument{
			"id_organization": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"id_category": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id_organization, ok := p.Args["id_organization"].(string)
			if !ok {
				id_organization = ""
			}
			id_category, ok := p.Args["id_category"].(string)
			if !ok {
				id_category = ""
			}

			return controller.catalogUseCase.GetProductList(p.Context, id_organization, id_category)
		},
	}
}

func (controller CatalogController) GetProductDetail() *graphql.Field {
	return &graphql.Field{
		Type: catalogfield.CategoryInfo,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if !ok {
				return nil, catalogpkg.ErrProductNotFound
			}
			return controller.catalogUseCase.GetProductDetail(p.Context, id)
		},
	}
}
