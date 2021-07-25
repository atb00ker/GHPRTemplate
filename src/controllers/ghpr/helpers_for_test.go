// +build integration_tests unit_tests all_tests

package ghpr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

// Set of type / funcs used by multiple tests in ghpr
var correctAuthToken = os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")

type mockHasuraClient struct{}

func (m *mockHasuraClient) insertGhpr(variables insertGhprArgs,
	auth []string) (response insertGhprOutput, err error) {

	value := auth[0]
	switch value {
	case "error":
		err = errors.New("Failed")
		return response, err
	}
	response = insertGhprOutput{Id: "default", Ghpr: &value}
	return response, err
}

func sendInsertGhprActionRequest(ghprValue string, authToken string,
	handlerFunc *http.HandlerFunc) *httptest.ResponseRecorder {

	var actionPayload = insertGhprPayload{
		SessionVariables: map[string]interface{}{},
		Input:            insertGhprArgs{Ghpr: ghprValue},
	}
	reqBytes, _ := json.Marshal(actionPayload)
	// InsertGhprAction
	return genericRequestSender(InsertGhprPath, bytes.NewBuffer(reqBytes), handlerFunc, authToken)
}

func genericRequestSender(path string, body io.Reader,
	handler *http.HandlerFunc, authToken string) *httptest.ResponseRecorder {

	req, _ := http.NewRequest("POST", path, body)
	req.Header.Set("Admin-Secret", authToken)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	return response
}
