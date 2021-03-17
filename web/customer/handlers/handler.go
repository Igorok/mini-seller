package handlers

import (
	"log"
	"mini-seller/domain/packages/customer/productpkg/productrepository"
	"mini-seller/domain/packages/customer/productpkg/productusecase"
	"mini-seller/web/customer/schemas"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetHandler - http handler for graphql
func GetHandler(db *mongo.Database) gin.HandlerFunc {
	// initialization of use cases
	productRepository := productrepository.NewProductRepository(db)
	productUseCase := productusecase.NewUseCase(productRepository)

	// initialization of web schema Graphql
	gqlSchema, err := schemas.GetSchema(productUseCase)
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
