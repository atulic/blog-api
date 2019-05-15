package gql

import (
	"graphql-api/postgres"

	"github.com/graphql-go/graphql"
)

// Resolver struct holds the connection to our postgres db
type Resolver struct {
	db *postgres.Db
}

// UserResolver resolves the query through a call to the db
func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the id from arguments and assert type
	name, ok := p.Args["id"].(int)
	if ok {
		posts := r.db.GetPostByID(name)
		return posts, nil
	}

	return nil, nil
}
