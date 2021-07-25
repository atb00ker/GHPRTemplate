package ghpr

import (
	"encoding/json"
	"ghpr/src/controllers/graphql"
	"net/http"
)

const (
	InsertGhprPath = "/InsertGhpr"
)

type GhprCaller struct {
	Hasura HasuraInterface
}

func (caller *GhprCaller) InsertGhprAction(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	requestBody := ensureValidBody(response, request)

	var actionPayload insertGhprPayload
	if err := json.Unmarshal(requestBody, &actionPayload); err != nil {
		http.Error(response, "Incorrect Variables Provided", http.StatusBadRequest)
		return
	}

	auth := request.Header["Admin-Secret"]
	result, err := caller.Hasura.insertGhpr(actionPayload.Input, auth)
	if err != nil {
		errorObject := graphql.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		response.WriteHeader(http.StatusBadRequest)
		response.Write(errorBody)
		return
	}

	data, _ := json.Marshal(result)
	response.Write(data)
}
