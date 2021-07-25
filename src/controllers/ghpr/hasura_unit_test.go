// +build unit_tests all_tests

package ghpr

import (
	"bytes"
	"encoding/json"
	"errors"
	"ghpr/src/controllers/graphql"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var testInsertGhpr_ResponseCode_Table = []struct {
	auth []string
}{
	{[]string{"http-error"}},
	{[]string{"request-error"}},
	{[]string{"response-error"}},
	{[]string{"success"}},
}

type mockHTTPClient struct{}

func (m *mockHTTPClient) Do(req *http.Request) (response *http.Response, err error) {
	auth := req.Header["x-hasura-admin-secret"]
	if auth[0] == "http-error" {
		return &http.Response{
			StatusCode: 400,
		}, errors.New("http-error")
	} else if auth[0] == "request-error" {
		data := insertGhprArgsResponse{
			Data:   insertGhprOutputData{},
			Errors: []graphql.GraphQLError{{Message: auth[0]}},
		}
		dataJson, _ := json.Marshal(data)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(dataJson))),
		}, nil
	} else if auth[0] == "response-error" {
		dataJson, _ := json.Marshal("{\"incorrect-data\": 10}")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(dataJson))),
		}, nil
	}
	data := insertGhprArgsResponse{
		Data: insertGhprOutputData{
			Insert_ghpr_one: insertGhprOutput{
				Id:   "success",
				Ghpr: &auth[0],
			},
		},
		Errors: []graphql.GraphQLError{},
	}
	dataJson, _ := json.Marshal(data)
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(dataJson))),
	}, nil
}

func TestInsertGhpr_ResponseCode(t *testing.T) {
	client := &mockHTTPClient{}
	for _, row := range testInsertGhpr_ResponseCode_Table {
		result, err := sendInsertGhprRequest(insertGhprArgs{Ghpr: ""}, row.auth, client)
		if row.auth[0] == "success" && (result.Id != "success" || *result.Ghpr != row.auth[0]) {
			t.Errorf("Expected: %s. Got: %s", row.auth[0], result.Id)
			return
		}
		error_message := err.Error()
		if row.auth[0] == "response-error" {
			if !strings.Contains(error_message, "json: cannot unmarshal") {
				t.Errorf("Expected: %s. Got: %s", row.auth[0], result.Id)
			}
			return
		}
		if error_message != row.auth[0] {
			t.Errorf("Error not thrown, instead got %s", error_message)
		}
	}
}
