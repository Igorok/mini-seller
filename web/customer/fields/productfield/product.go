package productfield

import "github.com/graphql-go/graphql"

var ProductList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProductList",
		Fields: graphql.Fields{
			"products": &graphql.Field{
				Type: graphql.NewList(ProductForList),
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ProductForList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProductForList",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
			"Category": &graphql.Field{
				Type: CategoryForProduct,
			},
			"Organization": &graphql.Field{
				Type: OrganizationForProduct,
			},
		},
	},
)

var CategoryForProduct = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CategoryForProduct",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var OrganizationForProduct = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "OrganizationForProduct",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
