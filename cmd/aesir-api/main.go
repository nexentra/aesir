package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/nexentra/aesir/internals/graph"
	"github.com/rs/cors"
)

func main() {
	port := flag.String("port", "8080", "port to run the server")
	flag.Parse()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            false,
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
