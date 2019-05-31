package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atulic/blog-api/gql"
	"github.com/atulic/blog-api/postgres"
	"github.com/atulic/blog-api/server"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	// Initialize our api and return a pointer to our router for http.ListenAndServe
	// and a pointer to our db to defer its closing when main() is finished
	router, db := initializeAPI()
	defer db.Close()

	// Listen on port 4000 and if there's an error log it and exit
	log.Fatal(http.ListenAndServe(":4000", router))
}

// openDbConnection creates a new connection to our postgres database
func openDbConnection() (db *postgres.DbConnection) {
	db, err := postgres.New(
		postgres.BuildDbConnString("db", 5432, "postgres", "password", "go_graphql_db"),
	)

	if err != nil {
		return nil
	}

	return db
}

func initializeAPI() (*chi.Mux, *postgres.DbConnection) {
	// Create a new router
	router := chi.NewRouter()

	db := openDbConnection()

	repository := postgres.NewPostgresRepository(db.DB)

	// Create our root query for graphql
	rootQuery := gql.NewRoot(repository)

	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: rootQuery.Mutation},
	)

	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Basic configuration for CORS middleware. Allows all origins for development purposes.
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// Add some middleware to our router
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		cors.Handler,               // handle cross-origin requests
		middleware.Logger,          // log api request calls
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}
