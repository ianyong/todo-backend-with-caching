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
)

type CreateTodoTestSuite struct {
	suite.Suite
	tests.TestComponents
}

func (s *CreateTodoTestSuite) SetupTest() {
	s.TestComponents = tests.SetUp()
}

const validJSON = `{
    "name": "Todo",
    "description": "This is a valid todo.",
    "dueDate": "2021-11-01T00:00:00.000000+00:00"
}`

var validTodo = todoviews.View{
	ID:          1,
	Name:        "Todo",
	Description: "This is a valid todo.",
	DueDate:     time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC).Local(),
}

func (s *CreateTodoTestSuite) TestCreateTodo() {
	request, err := http.NewRequest(http.MethodPost, "/api/v1/todos", strings.NewReader(validJSON))
	if err != nil {
		s.T().Errorf("Error creating request: %v", err)
	}

	responseRecorder := httptest.NewRecorder()
	s.Router.ServeHTTP(responseRecorder, request)

	tests.CheckResponseCode(s.T(), http.StatusCreated, responseRecorder.Code)

	body := tests.GetResponseBody(s.T(), responseRecorder.Body)
	var actualTodo todoviews.View
	err = json.Unmarshal(body.Payload, &actualTodo)
	if err != nil {
		s.T().Errorf("Error decoding response body: %v", err)
	}

	expectedTodo := validTodo

	tests.CheckResponseBody(s.T(), expectedTodo, actualTodo)

	err = s.TruncateTables("todos")
	if err != nil {
		s.T().Errorf("Error truncating tables: %v", err)
	}
}

const missingFieldJSON = `{
    "name": "Todo",
    "dueDate": "2021-11-01T00:00:00.000000+00:00"
}`

func (s *CreateTodoTestSuite) TestCreateTodoMissingField() {
	request, err := http.NewRequest(http.MethodPost, "/api/v1/todos", strings.NewReader(missingFieldJSON))
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

const extraFieldJSON = `{
    "id": 3,
    "name": "Todo",
    "description": "This is a valid todo.",
    "dueDate": "2021-11-01T00:00:00.000000+00:00"
}`

func (s *CreateTodoTestSuite) TestCreateTodoExtraField() {
	request, err := http.NewRequest(http.MethodPost, "/api/v1/todos", strings.NewReader(extraFieldJSON))
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

func TestCreateTodoTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTodoTestSuite))
}
