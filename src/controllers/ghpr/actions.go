package ghpr

import (
	"encoding/json"
	"ghpr/src/controllers/graphql"
	"net/http"
)

const (
	InsertGhprPath = "/InsertGhpr"
)

func InsertGhprAction(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	requestBody := ensureValidBody(response, request)

	var actionPayload insertGhprPayload
	var err = json.Unmarshal(requestBody, &actionPayload)
	if err != nil {
		http.Error(response, "Incorrect Variables Provided", http.StatusBadRequest)
		return
	}

	auth := request.Header["Admin-Secret"]
	result, err := insertGhpr(actionPayload.Input, auth)

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
