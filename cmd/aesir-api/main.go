package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nexentra/aesir/ent"
	"github.com/nexentra/aesir/graph"
	"github.com/nexentra/aesir/graph/generated"

	"github.com/rs/cors"
)



func newClient() *ent.Client {
	// Create a Sqlite3 database connection.
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func main() {
	port := flag.String("port", "8080", "port to run the server")
	flag.Parse()

	client := newClient()
	defer client.Close()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            false,
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))
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
