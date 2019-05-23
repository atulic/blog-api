package resolvers

import (
	"github.com/graphql-go/graphql"
)

// PostQueryResolver resolves the query through a call to the db
func (r *Resolver) PostQueryResolver(params graphql.ResolveParams) (interface{}, error) {
	// Strip the id from arguments and assert type
	name, ok := params.Args["id"].(int)
	if ok {
		posts := r.Db.GetPostByID(name)
		return posts, nil
	}

	return nil, nil
}
