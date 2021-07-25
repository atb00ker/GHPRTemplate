// +build integration_tests all_tests

package ghpr

import (
	"bytes"
	"encoding/json"
	"ghpr/src/controllers/graphql"
	"io/ioutil"
	"net/http"
	"testing"
)

func cleanDb() {
	jsonStr := []byte(`{"query":"mutation { delete_ghpr(where: {}) { affected_rows }}","variables":null}`)
	client := http.Client{}
	graphRequest, _ := http.NewRequest("POST", graphql.Endpoint, bytes.NewBuffer(jsonStr))
	graphRequest.Header = http.Header{
		"Content-Type":          []string{"application/json"},
		"x-hasura-admin-secret": []string{correctAuthToken},
	}
	client.Do(graphRequest)
}

func TestInsertGhpr(t *testing.T) {
	ghprValue := "test"
	testHasuraGhprCaller := GhprCaller{Hasura: &HasuraClient{}}
	handler := http.HandlerFunc(testHasuraGhprCaller.InsertGhprAction)
	response := sendInsertGhprActionRequest(ghprValue, correctAuthToken, &handler)

	// Assert HTTP status
	if response.Code != 200 {
		t.Errorf("Expected 200, got %d", response.Code)
	}
	respBytes, _ := ioutil.ReadAll(response.Body)
	var resp insertGhprOutput
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		t.Error("Got nil response!")
	}

	if resp.Id == "" {
		t.Error("Did not receive Id in hasura response")
	}
	if *resp.Ghpr != ghprValue {
		t.Errorf("Expected %s, got %s", ghprValue, *resp.Ghpr)
	}

	// Clean Database
	cleanDb()
}

var testServerAuth_AuthTable = []struct {
	in  string
	out int
}{
	{"", 400},
	{correctAuthToken, 200},
	{"someIncorrectValue", 400},
}

func TestServerAuth(t *testing.T) {
	testHasuraGhprCaller := GhprCaller{Hasura: &HasuraClient{}}
	handler := http.HandlerFunc(testHasuraGhprCaller.InsertGhprAction)
	for _, row := range testServerAuth_AuthTable {
		response := sendInsertGhprActionRequest("", row.in, &handler)
		if response.Code != row.out {
			t.Errorf("Expected %d, got %d", row.out, response.Code)
		}
	}
}
