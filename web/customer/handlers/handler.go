package handlers

import (
	"log"
	"mini-seller/domain/packages/customer/catalog/catalogrepository"
	"mini-seller/domain/packages/customer/catalog/catalogusecase"
	"mini-seller/web/customer/schemas"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetHandler - http handler for graphql
func GetHandler(db *mongo.Database) gin.HandlerFunc {
	// initialization of use cases
	catalogRepository := catalogrepository.NewCatalogRepository(db)
	catalogUseCase := catalogusecase.NewUseCase(catalogRepository)

	// initialization of web schema Graphql
	gqlSchema, err := schemas.GetSchema(catalogUseCase)
	if err != nil {
		log.Fatal("GetSchema:", err)
	}

	return func(c *gin.Context) {
		// Creates a GraphQL-go HTTP handler with the defined schema
		h := handler.New(&handler.Config{
			Schema:   &gqlSchema,
			Pretty:   true,
			GraphiQL: true,
		})

		h.ServeHTTP(c.Writer, c.Request)
	}
}
