package todohandlers

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

type DestroyTodoTestSuite struct {
	suite.Suite
	tests.TestComponents
}

func (s *DestroyTodoTestSuite) SetupTest() {
	s.TestComponents = tests.SetUp()
}

func (s *DestroyTodoTestSuite) TestDestroyTodo() {
	err := testseeds.SeedTodos(s.DB)
	if err != nil {
		s.T().Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest(http.MethodDelete, "/api/v1/todos/4", nil)
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

	expectedTodo := todoviews.View{}

	tests.CheckResponseBody(s.T(), expectedTodo, actualTodo)

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

func (s *DestroyTodoTestSuite) TestDestroyNonExistentTodo() {
	request, err := http.NewRequest(http.MethodDelete, "/api/v1/todos/4", nil)
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

func TestDestroyTodoTestSuite(t *testing.T) {
	suite.Run(t, new(DestroyTodoTestSuite))
}
