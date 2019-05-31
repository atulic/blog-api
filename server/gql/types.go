package gql

import "github.com/graphql-go/graphql"

// Post describes a graphql object containing a Post
var Post = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"posted": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
