package ghpr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ensureValidBody(response http.ResponseWriter, request *http.Request) []byte {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Invalid Payload", http.StatusBadRequest)
		return requestBody
	}
	return requestBody
}

func insertGhpr(variables insertGhprArgs, auth []string) (response insertGhprOutput, err error) {
	reqBody := insertGhprArgsRequest{
		Query:     "mutation ($ghpr: String!) {   insert_ghpr_one(object: {ghpr: $ghpr}) { Id, ghpr }}",
		Variables: variables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	client := http.Client{}
	graphRequest, err := http.NewRequest("POST", "http://localhost:8080/v1/graphql", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	graphRequest.Header = http.Header{
		"Content-Type":          []string{"application/json"},
		"x-hasura-admin-secret": auth,
	}

	resp, err := client.Do(graphRequest)
	if err != nil {
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var hasuraResponse insertGhprArgsResponse
	err = json.Unmarshal(respBytes, &hasuraResponse)
	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_ghpr_one
	return
}
