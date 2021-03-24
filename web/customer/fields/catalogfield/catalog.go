package catalogfield

import "github.com/graphql-go/graphql"

var OrganizationInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "OrganizationInfo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"categories": &graphql.Field{
				Type: graphql.NewList(CategoryInfo),
			},
			"products": &graphql.Field{
				Type: graphql.NewList(ProductInfo),
			},
		},
	},
)

var CategoryInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CategoryInfo",
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
		},
	},
)

var ProductInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProductInfo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"idCategory": &graphql.Field{
				Type: graphql.String,
			},
			"idOrganization": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
