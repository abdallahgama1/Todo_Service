package handlers_test

import (
	"Todo_Service/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func registerAndLogin(t *testing.T, username, password string) string {
	registerInput := models.RegisterInput{Username: username, Password: password}
	body, _ := json.Marshal(registerInput)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	loginInput := models.LoginInput{Username: username, Password: password}
	body, _ = json.Marshal(loginInput)
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	return response["token"]
}

func TestCreateAndGetTodos(t *testing.T) {
	token := registerAndLogin(t, "todotestuser", "password")

	todoInput := models.CreateTodoRequest{Title: "My Test Todo", Category: "Testing", Priority: "High"}
	body, _ := json.Marshal(todoInput)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var createdTodo models.Todo
	json.Unmarshal(w.Body.Bytes(), &createdTodo)
	assert.Equal(t, "My Test Todo", createdTodo.Title)

	url := fmt.Sprintf("/todos/%d", createdTodo.ID)
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, _ = http.NewRequest("GET", "/todos", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var todos []models.Todo
	json.Unmarshal(w.Body.Bytes(), &todos)
	assert.Len(t, todos, 1)
	assert.Equal(t, "My Test Todo", todos[0].Title)
}

func TestTodoPermissions(t *testing.T) {
	tokenUserA := registerAndLogin(t, "userA", "password")
	tokenUserB := registerAndLogin(t, "userB", "password")

	todoInput := models.CreateTodoRequest{Title: "User A's Todo", Category: "A", Priority: "Low"}
	body, _ := json.Marshal(todoInput)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+tokenUserA)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var createdTodo models.Todo
	json.Unmarshal(w.Body.Bytes(), &createdTodo)

	url := fmt.Sprintf("/todos/%d", createdTodo.ID)
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+tokenUserB)
	w = httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
