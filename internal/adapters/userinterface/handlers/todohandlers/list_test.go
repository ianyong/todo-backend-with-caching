package todohandlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/tests"
	"github.com/ianyong/todo-backend/internal/tests/testseeds"
)

var testComponents *tests.TestComponents

func TestMain(m *testing.M) {
	testComponents = tests.SetUp()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestListTodos(t *testing.T) {
	err := testseeds.SeedTodos(testComponents.DB)
	if err != nil {
		t.Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := api.WrapHandler(testComponents.Services, List)
	handler.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(t, http.StatusOK, responseRecorder.Code)

	body := tests.GetResponseBody(t, responseRecorder.Body)
	var actualTodos []todoviews.ListView
	err = json.Unmarshal(body.Payload, &actualTodos)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedTodos := make([]todoviews.ListView, len(testseeds.TodoSeeds))
	for i := range testseeds.TodoSeeds {
		todo := testseeds.TodoSeeds[i]
		expectedTodos[i] = todoviews.ListViewFrom(&todo)
	}

	tests.CheckResponseBody(t, expectedTodos, actualTodos)

	err = testComponents.TruncateTables("todos")
	if err != nil {
		t.Errorf("Error truncating tables: %v", err)
	}
}

func TestListTodosEmptyCollection(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := api.WrapHandler(testComponents.Services, List)
	handler.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(t, http.StatusOK, responseRecorder.Code)

	body := tests.GetResponseBody(t, responseRecorder.Body)
	var actualTodos []todoviews.ListView
	err = json.Unmarshal(body.Payload, &actualTodos)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedTodos := []todoviews.ListView{}

	tests.CheckResponseBody(t, expectedTodos, actualTodos)
}