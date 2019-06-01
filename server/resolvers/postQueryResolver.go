package resolvers

import (
	"github.com/graphql-go/graphql"
)

// PostQueryResolver resolves the query through a call to the db
func (r *Resolver) PostQueryResolver(params graphql.ResolveParams) (interface{}, error) {
	// Strip the id from arguments and assert type
	id, ok := params.Args["id"].(int)
	if ok {
		posts, err := r.Repository.GetByID(id)
		return posts, err
	}

	// We didn't get a valid ID as a param, so we return all the posts
	return r.Repository.GetAllPosts()
}
