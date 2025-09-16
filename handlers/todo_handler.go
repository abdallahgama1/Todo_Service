package handlers

import (
	"Todo_Service/models"
	"Todo_Service/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func getUserId(c *gin.Context) (uint, bool) {
	val, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return 0, false
	}
	userId, ok := val.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type in context"})
		return 0, false
	}
	return userId, true
}

// CreateTodo godoc
// @Summary      Create a new todo
// @Description  Adds a new todo item to the user's list
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      models.CreateTodoRequest  true  "Todo Info"
// @Success      200   {object}  models.Todo
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req models.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, ok := getUserId(c)
	if !ok {
		return
	}

	todo, err := h.service.CreateTodo(req, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// GetAllTodos godoc
// @Summary      Get all todos for the current user
// @Description  Retrieves a list of all todo items belonging to the authenticated user
// @Tags         todos
// @Produce      json
// @Success      200  {array}   models.Todo
// @Failure      401  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /todos [get]
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	userId, ok := getUserId(c)
	if !ok {
		return
	}

	todos, err := h.service.GetAllTodos(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
// @Summary      Get a single todo by its ID
// @Description  Get details of a specific todo owned by the user
// @Tags         todos
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  models.Todo
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	userId, ok := getUserId(c)
	if !ok {
		return
	}

	todo, err := h.service.GetTodoByID(uint(id), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
// @Summary      Update an existing todo
// @Description  Update the details of a todo item owned by the user
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      int                      true  "Todo ID"
// @Param        todo  body      models.UpdateTodoRequest true  "Updated Todo Info"
// @Success      200   {object}  models.Todo
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req models.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, ok := getUserId(c)
	if !ok {
		return
	}

	todo, err := h.service.UpdateTodo(uint(id), req, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or failed to update"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Description  Deletes a todo item owned by the user
// @Tags         todos
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	userId, ok := getUserId(c)
	if !ok {
		return
	}

	if err := h.service.DeleteTodo(uint(id), userId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
