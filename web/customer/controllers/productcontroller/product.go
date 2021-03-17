package productcontroller

import (
	"mini-seller/domain/packages/customer/productpkg"
	"mini-seller/web/customer/fields/productfield"

	"github.com/graphql-go/graphql"
)

// ProductController - web controller
type ProductController struct {
	productUseCase productpkg.IUseCase
}

// NewProductController - constructor
func NewProductController(productUseCase productpkg.IUseCase) *ProductController {
	return &ProductController{productUseCase: productUseCase}
}

// Products - products list
func (controller ProductController) GetProductList() *graphql.Field {
	return &graphql.Field{
		Type: productfield.ProductList,
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
				return nil, productpkg.ErrListLimit
			}
			limit, ok := p.Args["limit"].(int)
			if !ok {
				return nil, productpkg.ErrListLimit
			}

			return controller.productUseCase.GetProductList(p.Context, skip, limit)
		},
	}
}
