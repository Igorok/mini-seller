package main

import (
	"fmt"
	"log"
	"mini-seller/application/graph"
	"mini-seller/application/graph/generated"
	"mini-seller/domain/packages/catalogpkg/catalogrepository"
	"mini-seller/domain/packages/catalogpkg/catalogusecase"
	"mini-seller/infrastructure/mongohelper"
	"mini-seller/infrastructure/viperhelper"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/viper"
)

const defaultPort = "8080"

func main() {
	vip := viperhelper.Viper{ConfigType: "", ConfigName: "", ConfigPath: "infrastructure/viperhelper"}
	vip.Read()

	fmt.Println("WEB_PORT", viper.GetString("WEB_PORT"))

	db, err := mongohelper.Connect("")
	if err != nil {
		log.Fatal(err)
	}

	catalogRepository := catalogrepository.NewCatalogRepository(db)
	catalogUseCase := catalogusecase.NewCatalogUseCase(catalogRepository)

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	port := viper.GetString("WEB_PORT")

	resolver := graph.Resolver{CatalogUseCase: catalogUseCase}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
