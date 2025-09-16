package services

import (
	"Todo_Service/models"
	"Todo_Service/repositories"
	"strings"
	"time"

	"gorm.io/gorm"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(req models.CreateTodoRequest, userId uint) (*models.Todo, error) {
	newTodo := &models.Todo{
		Title:     req.Title,
		Completed: req.Completed,
		Category:  req.Category,
		Priority:  strings.Title(strings.ToLower(req.Priority)),
		UserID:    userId,
	}

	if req.Completed {
		now := gorm.DeletedAt{Time: time.Now().UTC(), Valid: true}
		newTodo.CompletedAt = &now
	}

	if req.DueDate != nil {
		parsedTime, _ := time.Parse(time.RFC3339, *req.DueDate)
		dueDate := gorm.DeletedAt{Time: parsedTime, Valid: true}
		newTodo.DueDate = &dueDate
	}

	return s.repo.Create(newTodo)
}

func (s *TodoService) UpdateTodo(id uint, req models.UpdateTodoRequest, userId uint) (*models.Todo, error) {
	existingTodo, err := s.repo.GetByIDForUser(id, userId)
	if err != nil {
		return nil, err
	}

	existingTodo.Title = req.Title
	existingTodo.Completed = req.Completed
	existingTodo.Category = req.Category
	existingTodo.Priority = strings.Title(strings.ToLower(req.Priority))

	if req.Completed && existingTodo.CompletedAt == nil {
		now := gorm.DeletedAt{Time: time.Now().UTC(), Valid: true}
		existingTodo.CompletedAt = &now
	} else if !req.Completed {
		existingTodo.CompletedAt = nil
	}

	if req.DueDate != nil {
		parsedTime, _ := time.Parse(time.RFC3339, *req.DueDate)
		dueDate := gorm.DeletedAt{Time: parsedTime, Valid: true}
		existingTodo.DueDate = &dueDate
	} else {
		existingTodo.DueDate = nil
	}

	return s.repo.Update(existingTodo)
}

func (s *TodoService) DeleteTodo(id, userId uint) error {
	todo, err := s.repo.GetByIDForUser(id, userId)
	if err != nil {
		return err
	}
	return s.repo.Delete(todo)
}

func (s *TodoService) GetAllTodos(userId uint) ([]models.Todo, error) {
	return s.repo.GetAllForUser(userId)
}

func (s *TodoService) GetTodoByID(id, userId uint) (*models.Todo, error) {
	return s.repo.GetByIDForUser(id, userId)
}

func (s *TodoService) GetTodosByCategory(category string, userId uint) ([]models.Todo, error) {
	return s.repo.GetByCategoryForUser(category, userId)
}

func (s *TodoService) GetTodosByStatus(completed bool, userId uint) ([]models.Todo, error) {
	return s.repo.GetByStatusForUser(completed, userId)
}

func (s *TodoService) SearchTodosByTitle(query string, userId uint) ([]models.Todo, error) {
	return s.repo.SearchByTitleForUser(query, userId)
}

func (s *TodoService) UpdateStatusByCategory(category string, completed bool, userId uint) ([]models.Todo, error) {
	var completedAt *gorm.DeletedAt
	if completed {
		now := gorm.DeletedAt{Time: time.Now().UTC(), Valid: true}
		completedAt = &now
	}

	err := s.repo.UpdateStatusByCategoryForUser(category, completed, completedAt, userId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByCategoryForUser(category, userId)
}

func (s *TodoService) DeleteAllTodos(userId uint) error {
	return s.repo.DeleteAllForUser(userId)
}
