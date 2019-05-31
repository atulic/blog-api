package resolvers

import (
	"time"

	"github.com/atulic/blog-api/postgres"
	"github.com/graphql-go/graphql"
)

// PostCreationResolver resolves the query and creates the row in the db
func (r *Resolver) PostCreationResolver(params graphql.ResolveParams) (interface{}, error) {
	post := postgres.Post{
		Title:   params.Args["title"].(string),
		Content: params.Args["content"].(string),
		Posted:  time.Now(),
	}

	r.Repository.Create(post)

	return post, nil
}
