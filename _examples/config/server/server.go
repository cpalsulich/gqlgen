package main

import (
	"log"
	"net/http"

	todo "github.com/cpalsulich/gqlgen/_examples/config"
	"github.com/cpalsulich/gqlgen/graphql/handler"
	"github.com/cpalsulich/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
