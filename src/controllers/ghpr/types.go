package ghpr

import "ghpr/src/controllers/graphql"

type Uuid string

type insertGhprOutput struct {
	Id   Uuid    `json:"Id"`
	Ghpr *string `json:"ghpr"`
}

type Mutation struct {
	GhprAction *insertGhprOutput
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
