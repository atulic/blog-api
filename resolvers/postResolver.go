package resolvers

import (
	"time"

	"github.com/atulic/blog-api/postgres"

	"github.com/graphql-go/graphql"
)

// Resolver struct holds the connection to our postgres db
type Resolver struct {
	Db *postgres.Db
}

// PostResolver resolves the query through a call to the db
func (r *Resolver) PostResolver(params graphql.ResolveParams) (interface{}, error) {
	// Strip the id from arguments and assert type
	name, ok := params.Args["id"].(int)
	if ok {
		posts := r.Db.GetPostByID(name)
		return posts, nil
	}

	return nil, nil
}

// PostCreationResolver resolves the query and creates the row in the db
func (r *Resolver) PostCreationResolver(params graphql.ResolveParams) (interface {}, error) {
	post := postgres.Post{
		ID: 2,
		Title: params.Args["title"].(string),
		Content: params.Args["content"].(string),
		Posted: time.Now(),
	}

	r.Db.CreatePost(post)

	return post, nil
}