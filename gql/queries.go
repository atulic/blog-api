package gql

import (
	"graphql-api/postgres"

	"github.com/graphql-go/graphql"
)

// Root struct holds a pointer to our graphQL object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"posts": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(Post),
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}
	return &root
}
