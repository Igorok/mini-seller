package catalogfield

import "github.com/graphql-go/graphql"

// graphql.NewList(catalogfield.ProductForCatalog),
var Catalog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Catalog",
		Fields: graphql.Fields{
			"products": &graphql.Field{
				Type: graphql.NewList(ProductForCatalog),
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ProductForCatalog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProductForCatalog",
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
				Type: CategoryForCatalog,
			},
			"Organization": &graphql.Field{
				Type: OrganizationForCatalog,
			},
		},
	},
)

var CategoryForCatalog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CategoryForCatalog",
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

var OrganizationForCatalog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "OrganizationForCatalog",
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
