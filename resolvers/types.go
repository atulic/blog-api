package resolvers

import (
	"github.com/atulic/blog-api/postgres"
)

// Resolver struct holds the connection to our postgres db
type Resolver struct {
	Repository postgres.Repository
}
