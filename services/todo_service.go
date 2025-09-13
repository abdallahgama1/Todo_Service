package services

import (
	"Todo_Service/models"
	"Todo_Service/repositories"
	"Todo_Service/utils"
	"database/sql"
	"strings"
	"time"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(req models.CreateTodoRequest) (*models.Todo, error) {
	if err := utils.ValidatePriority(req.Priority); err != nil {
		return nil, err
	}
	if req.DueDate != nil {
		if err := utils.ValidateDueDate(*req.DueDate); err != nil {
			return nil, err
		}
	}

	var dueDate sql.NullString
	if req.DueDate != nil {
		dueDate.String = *req.DueDate
		dueDate.Valid = true
	}

	var completedAt sql.NullString
	if req.Completed {
		completedAt.String = time.Now().UTC().Format(time.RFC3339)
		completedAt.Valid = true
	}

	newTodo := models.Todo{
		Title:       req.Title,
		Completed:   req.Completed,
		Category:    req.Category,
		Priority:    strings.Title(strings.ToLower(req.Priority)),
		DueDate:     dueDate,
		CompletedAt: completedAt,
	}

	return s.repo.Create(newTodo)
}

func (s *TodoService) UpdateTodo(id int, req models.UpdateTodoRequest) (*models.Todo, error) {
	if err := utils.ValidatePriority(req.Priority); err != nil {
		return nil, err
	}
	if req.DueDate != nil {
		if err := utils.ValidateDueDate(*req.DueDate); err != nil {
			return nil, err
		}
	}

	if _, err := s.repo.GetByID(id); err != nil {
		return nil, err
	}

	var dueDate sql.NullString
	if req.DueDate != nil {
		dueDate.String = *req.DueDate
		dueDate.Valid = true
	}

	var completedAt sql.NullString
	if req.Completed {
		completedAt.String = time.Now().UTC().Format(time.RFC3339)
		completedAt.Valid = true
	}

	updatedTodo := models.Todo{
		Title:       req.Title,
		Completed:   req.Completed,
		Category:    req.Category,
		Priority:    strings.Title(strings.ToLower(req.Priority)),
		DueDate:     dueDate,
		CompletedAt: completedAt,
	}

	return s.repo.Update(id, updatedTodo)
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}
func (s *TodoService) GetTodoByID(id int) (*models.Todo, error) {
	return s.repo.GetByID(id)
}
func (s *TodoService) GetTodosByCategory(category string) ([]models.Todo, error) {
	return s.repo.GetByCategory(category)
}
func (s *TodoService) GetTodosByStatus(completed bool) ([]models.Todo, error) {
	return s.repo.GetByStatus(completed)
}
func (s *TodoService) SearchTodosByTitle(query string) ([]models.Todo, error) {
	return s.repo.SearchByTitle(query)
}
func (s *TodoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}
func (s *TodoService) DeleteAllTodos() error {
	return s.repo.DeleteAll()
}

func (s *TodoService) UpdateStatusByCategory(category string, completed bool) ([]models.Todo, error) {
	var completedAt sql.NullString
	if completed {
		completedAt.String = time.Now().UTC().Format(time.RFC3339)
		completedAt.Valid = true
	}

	err := s.repo.UpdateStatusByCategory(category, completed, completedAt)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByCategory(category)
}
