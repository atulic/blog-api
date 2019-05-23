package resolvers

import (
	"time"

	"github.com/atulic/blog-api/postgres"

	"github.com/graphql-go/graphql"
)

// PostUpdateResolver resolves the query and creates the row in the db
func (r *Resolver) PostUpdateResolver(params graphql.ResolveParams) (interface{}, error) {
	post := postgres.Post{
		ID:      params.Args["id"].(int),
		Title:   params.Args["title"].(string),
		Content: params.Args["content"].(string),
		Posted:  time.Now(),
	}

	r.Db.UpdatePost(post)

	return post, nil
}
