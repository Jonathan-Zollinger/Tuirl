package main

import (
	"os"

	"github.com/Jonathan-Zollinger/Tuirl/internal/cmd"

	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Jonathan-Zollinger/Tuirl/graph"
)

var version = "dev"

const defaultPort = "7080"

func main() {
	graphPlayground()
	cmd.Execute(
		version,
		os.Exit,
		os.Args[1:],
	)
}

func graphPlayground() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
