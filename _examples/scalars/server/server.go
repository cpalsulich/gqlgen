package main

import (
	"log"
	"net/http"

	"github.com/cpalsulich/gqlgen/_examples/scalars"
	"github.com/cpalsulich/gqlgen/graphql/handler"
	"github.com/cpalsulich/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
