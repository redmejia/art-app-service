package graphql

import "github.com/graphql-go/graphql"

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: query,
		},
	)

	return result
}
