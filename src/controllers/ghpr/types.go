package ghpr

import (
	"ghpr/src/controllers/graphql"
	"net/http"
)

type insertGhprOutput struct {
	Id   graphql.Uuid `json:"Id"`
	Ghpr *string      `json:"ghpr"`
}

type insertGhprArgs struct {
	Ghpr string `json:"ghpr"`
}

type insertGhprPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            insertGhprArgs         `json:"input"`
}

type insertGhprArgsRequest struct {
	Query     string         `json:"query"`
	Variables insertGhprArgs `json:"variables"`
}

type insertGhprOutputData struct {
	Insert_ghpr_one insertGhprOutput `json:"insert_ghpr_one"`
}

type insertGhprArgsResponse struct {
	Data   insertGhprOutputData   `json:"data,omitempty"`
	Errors []graphql.GraphQLError `json:"errors,omitempty"`
}

// Hasura
type HasuraInterface interface {
	insertGhpr(insertGhprArgs, []string) (insertGhprOutput, error)
}

type HasuraClient struct{}

// HTTP Client
type httpInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
