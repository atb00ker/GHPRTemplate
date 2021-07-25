package graphql

import "os"

// GraphQL has datatype UUID
type Uuid string

type GraphQLError struct {
	Message string `json:"message"`
}

var Endpoint = os.Getenv("HASURA_GRAPHQL_ENDPOINT") + os.Getenv("HASURA_GRAPHQL_API_PATHS_GRAPHQL")
