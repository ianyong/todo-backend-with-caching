package tests

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
)

func CheckResponseCode(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("Incorrect response code. Expected: %d. Got: %d.", expected, actual)
	}
}

func GetResponseBody(t *testing.T, body *bytes.Buffer) *api.Response {
	var response api.Response
	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		t.Errorf("Error when decoding response body: %v", err)
	}
	return &response
}

func CheckResponseBody(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Incorrect response body. Expected: %v. Got: %v.", expected, actual)
	}
}
