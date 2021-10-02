package todohandlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/tests"
	"github.com/ianyong/todo-backend/internal/tests/testseeds"
)

type UpdateTodoTestSuite struct {
	suite.Suite
	tests.TestComponents
}

func (s *UpdateTodoTestSuite) SetupTest() {
	s.TestComponents = tests.SetUp()
}

const validUpdateJSON = `{
    "id": 2,
    "name": "Todo",
    "description": "This is a valid todo.",
    "dueDate": "2021-11-01T00:00:00.000000+00:00",
    "isCompleted": true
}`

var updatedTodo = todoviews.View{
	ID:          2,
	Name:        "Todo",
	Description: "This is a valid todo.",
	DueDate:     time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC).Local(),
	IsCompleted: true,
}

func (s *CreateTodoTestSuite) TestUpdateTodo() {
	err := testseeds.SeedTodos(s.DB)
	if err != nil {
		s.T().Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest(http.MethodPut, "/api/v1/todos/2", strings.NewReader(validUpdateJSON))
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

	expectedTodo := updatedTodo

	tests.CheckResponseBody(s.T(), expectedTodo, actualTodo)

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

func (s *CreateTodoTestSuite) TestUpdateTodoIDMismatch() {
	err := testseeds.SeedTodos(s.DB)
	if err != nil {
		s.T().Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest(http.MethodPut, "/api/v1/todos/3", strings.NewReader(validUpdateJSON))
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

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

const missingFieldUpdateJSON = `{
    "id": 2,
    "name": "Todo",
    "description": "This is a valid todo.",
    "dueDate": "2021-11-01T00:00:00.000000+00:00"
}`

func (s *CreateTodoTestSuite) TestUpdateTodoMissingField() {
	err := testseeds.SeedTodos(s.DB)
	if err != nil {
		s.T().Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest(http.MethodPut, "/api/v1/todos/2", strings.NewReader(missingFieldUpdateJSON))
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

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

const extraFieldUpdateJSON = `{
    "id": 2,
    "name": "Todo",
    "hello": "world",
    "description": "This is a valid todo.",
    "dueDate": "2021-11-01T00:00:00.000000+00:00",
    "isCompleted": true
}`

func (s *CreateTodoTestSuite) TestUpdateTodoExtraField() {
	err := testseeds.SeedTodos(s.DB)
	if err != nil {
		s.T().Errorf("Error seeding todos: %v", err)
	}

	request, err := http.NewRequest(http.MethodPut, "/api/v1/todos/2", strings.NewReader(extraFieldUpdateJSON))
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

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

func TestUpdateTodoTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateTodoTestSuite))
}
