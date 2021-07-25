package ghpr

import (
	"bytes"
	"encoding/json"
	"errors"
	"ghpr/src/controllers/graphql"
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

func (c *HasuraClient) insertGhpr(variables insertGhprArgs,
	auth []string) (response insertGhprOutput, err error) {
	client := &http.Client{}
	return sendInsertGhprRequest(variables, auth, client)
}

func sendInsertGhprRequest(variables insertGhprArgs, auth []string,
	client httpInterface) (response insertGhprOutput, err error) {

	reqBody := insertGhprArgsRequest{
		Query:     "mutation ($ghpr: String!) {   insert_ghpr_one(object: {ghpr: $ghpr}) { Id, ghpr }}",
		Variables: variables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	graphRequest, err := http.NewRequest("POST", graphql.Endpoint, bytes.NewBuffer(reqBytes))
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
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var hasuraResponse insertGhprArgsResponse
	if err = json.Unmarshal(respBytes, &hasuraResponse); err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_ghpr_one
	return
}
