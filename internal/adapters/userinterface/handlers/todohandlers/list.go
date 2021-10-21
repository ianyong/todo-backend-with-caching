package todohandlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/cache/v8"

	"github.com/ianyong/todo-backend/internal/adapters/userinterface/api"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/json"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/views/todoviews"
	"github.com/ianyong/todo-backend/internal/services"
)

const cacheKey = "listTodos"

func List(r *http.Request, s *services.Services) (*api.Response, error) {
	var apiResponse *api.Response
	err := s.Cache.Get(s.CacheCtx, cacheKey, &apiResponse)
	if err == nil && apiResponse != nil {
		log.Println("Cache hit!")
		return apiResponse, nil
	}
	log.Println("Cache miss!")

	todos, err := s.TodoService.GetAllTodos()
	if err != nil {
		return nil, err
	}

	todoListViews := make([]todoviews.ListView, len(todos))
	for i := range todos {
		todo := todos[i]
		todoListViews[i] = todoviews.ListViewFrom(&todo)
	}

	data, err := json.EncodeView(todoListViews)
	if err != nil {
		return nil, err
	}

	apiResponse = &api.Response{
		Payload: data,
		Code:    http.StatusOK,
	}

	err = s.Cache.Set(&cache.Item{
		Ctx:   s.CacheCtx,
		Key:   cacheKey,
		Value: apiResponse,
		TTL:   10 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to save response to cache: %v\n", err)
	}

	return apiResponse, nil
}
