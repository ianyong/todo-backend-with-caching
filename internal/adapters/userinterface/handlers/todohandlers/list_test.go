package todohandlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/tests"
	"github.com/ianyong/todo-backend/internal/tests/testseeds"
)

type ListTodosTestSuite struct {
	suite.Suite
	tests.TestComponents
}

func (s *ListTodosTestSuite) SetupTest() {
	s.TestComponents = tests.SetUp()
}

func (s *ListTodosTestSuite) TestListTodos() {
	err := testseeds.SeedTodos(s.DB)
	if err != nil {
		s.T().Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		s.T().Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := api.WrapHandler(s.Services, List)
	handler.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(s.T(), http.StatusOK, responseRecorder.Code)

	body := tests.GetResponseBody(s.T(), responseRecorder.Body)
	var actualTodos []todoviews.ListView
	err = json.Unmarshal(body.Payload, &actualTodos)
	if err != nil {
		s.T().Errorf("Error decoding response body: %v", err)
	}

	expectedTodos := make([]todoviews.ListView, len(testseeds.TodoSeeds))
	for i := range testseeds.TodoSeeds {
		todo := testseeds.TodoSeeds[i]
		expectedTodos[i] = todoviews.ListViewFrom(&todo)
	}

	tests.CheckResponseBody(s.T(), expectedTodos, actualTodos)

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

func (s *ListTodosTestSuite) TestListTodosEmptyCollection() {
	request, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		s.T().Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := api.WrapHandler(s.Services, List)
	handler.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(s.T(), http.StatusOK, responseRecorder.Code)

	body := tests.GetResponseBody(s.T(), responseRecorder.Body)
	var actualTodos []todoviews.ListView
	err = json.Unmarshal(body.Payload, &actualTodos)
	if err != nil {
		s.T().Errorf("Error decoding response body: %v", err)
	}

	expectedTodos := []todoviews.ListView{}

	tests.CheckResponseBody(s.T(), expectedTodos, actualTodos)
}

func TestListTodosTestSuite(t *testing.T) {
	suite.Run(t, new(ListTodosTestSuite))
}
