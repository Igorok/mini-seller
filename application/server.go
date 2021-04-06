package main

import (
	"log"
	"mini-seller/application/graph/generated"
	"mini-seller/application/graph/resolvers"
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

	db, err := mongohelper.Connect("")
	if err != nil {
		log.Fatal(err)
	}

	catalogRepository := catalogrepository.NewCatalogRepository(db)
	catalogUseCase := catalogusecase.NewCatalogUseCase(catalogRepository)

	resolver := resolvers.Resolver{CatalogUseCase: catalogUseCase}

	router := http.NewServeMux()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver})))

	port := viper.GetString("WEB_PORT")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, resolvers.LoaderMiddleware(catalogUseCase, router)))
}
