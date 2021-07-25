// +build unit_tests all_tests

package ghpr

import (
	"bytes"
	"net/http"
	"testing"
)

var testInsertGhprAction_ResponseCode_Table = []struct {
	in   string
	out  int
	auth string
}{
	{"", 400, ""},
	{"{}", 200, ""},
	{"{}", 400, "error"},
}

func TestInsertGhprAction_ResponseCode(t *testing.T) {
	testActionGhprCaller := GhprCaller{Hasura: &mockHasuraClient{}}
	handler := http.HandlerFunc(testActionGhprCaller.InsertGhprAction)
	for _, row := range testInsertGhprAction_ResponseCode_Table {
		response := genericRequestSender(InsertGhprPath, bytes.NewBuffer([]byte(row.in)), &handler, row.auth)
		// Assert HTTP status
		if response.Code != row.out {
			t.Errorf("Expected %d, got %d", row.out, response.Code)
		}
	}
}
