package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// ExecuteQuery builds and runs the GraphQL queries
func ExecuteQuery(query string, variables map[string]interface{}, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: variables,
	})

	// Error check
	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}
