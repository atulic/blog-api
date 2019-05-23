package resolvers

import (
	"github.com/graphql-go/graphql"
)

// PostDeleteResolver resolves the query and creates the row in the db
func (r *Resolver) PostDeleteResolver(params graphql.ResolveParams) (interface{}, error) {

	id := params.Args["id"].(int)

	r.Db.DeletePost(id)

	return id, nil
}
