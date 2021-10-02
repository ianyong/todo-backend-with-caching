package todohandlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/tests"
	"github.com/ianyong/todo-backend/internal/tests/testseeds"
)

type ReadTodoTestSuite struct {
	suite.Suite
	tests.TestComponents
}

func (s *ReadTodoTestSuite) SetupTest() {
	s.TestComponents = tests.SetUp()
}

func (s *ReadTodoTestSuite) TestReadTodo() {
	err := testseeds.SeedTodos(s.DB)
	if err != nil {
		s.T().Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest(http.MethodGet, "/api/v1/todos/2", nil)
	if err != nil {
		s.T().Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	s.Router.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(s.T(), http.StatusOK, responseRecorder.Code)

	body := tests.GetResponseBody(s.T(), responseRecorder.Body)
	var actualTodo todoviews.View
	err = json.Unmarshal(body.Payload, &actualTodo)
	if err != nil {
		s.T().Errorf("Error decoding response body: %v", err)
	}

	seedTodo := testseeds.TodoSeeds[1]
	expectedTodo := todoviews.ViewFrom(&seedTodo)

	tests.CheckResponseBody(s.T(), expectedTodo, actualTodo)

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

func (s *ReadTodoTestSuite) TestReadNonExistentTodo() {
	request, err := http.NewRequest(http.MethodGet, "/api/v1/todos/2", nil)
	if err != nil {
		s.T().Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	s.Router.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(s.T(), http.StatusNotFound, responseRecorder.Code)

	body := tests.GetResponseBody(s.T(), responseRecorder.Body)
	var actualTodo todoviews.View
	err = json.Unmarshal(body.Payload, &actualTodo)
	if err != nil {
		s.T().Errorf("Error decoding response body: %v", err)
	}

	expectedTodo := todoviews.View{}

	tests.CheckResponseBody(s.T(), expectedTodo, actualTodo)
}

func (s *ReadTodoTestSuite) TestReadInvalidID() {
	request, err := http.NewRequest(http.MethodGet, "/api/v1/todos/abc", nil)
	if err != nil {
		s.T().Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	s.Router.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(s.T(), http.StatusBadRequest, responseRecorder.Code)

	body := tests.GetResponseBody(s.T(), responseRecorder.Body)
	var actualTodo todoviews.View
	err = json.Unmarshal(body.Payload, &actualTodo)
	if err != nil {
		s.T().Errorf("Error decoding response body: %v", err)
	}

	expectedTodo := todoviews.View{}

	tests.CheckResponseBody(s.T(), expectedTodo, actualTodo)
}

func TestReadTodoTestSuite(t *testing.T) {
	suite.Run(t, new(ReadTodoTestSuite))
}
