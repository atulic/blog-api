package gql

import (
	"github.com/atulic/blog-api/postgres"
	"github.com/atulic/blog-api/resolvers"

	"github.com/graphql-go/graphql"
)

// Root struct holds a pointer to our graphQL object
type Root struct {
	Query *graphql.Object
	Mutation *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	resolver := resolvers.Resolver{Db: db}

	queryType := graphql.NewObject(
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
					Resolve: resolver.PostResolver,
				},
			},
		},
	)

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields {
			"create": &graphql.Field{
				Type: Post,
				Description: "Create new post",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
			},
			Resolve: resolver.PostCreationResolver,
			},
		},
	})

	return &Root {
		Query: queryType,
		Mutation: mutationType,
	}
}
