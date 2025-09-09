package services

import (
	"Todo_Service/models"
	"Todo_Service/repositories"
	"Todo_Service/utils"
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

func (s *TodoService) CreateTodo(req models.CreateTodoRequest) (*models.Todo, error) {
	if err := utils.ValidatePriority(req.Priority); err != nil {
		return nil, err
	}
	if req.DueDate != nil {
		if err := utils.ValidateDueDate(*req.DueDate); err != nil {
			return nil, err
		}
	}

	newTodo := &models.Todo{
		Title:     req.Title,
		Completed: req.Completed,
		Category:  req.Category,
		Priority:  strings.Title(strings.ToLower(req.Priority)),
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

func (s *TodoService) UpdateTodo(id uint, req models.UpdateTodoRequest) (*models.Todo, error) {
	existingTodo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if err := utils.ValidatePriority(req.Priority); err != nil {
		return nil, err
	}
	if req.DueDate != nil {
		if err := utils.ValidateDueDate(*req.DueDate); err != nil {
			return nil, err
		}
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

func (s *TodoService) UpdateStatusByCategory(category string, completed bool) ([]models.Todo, error) {
	var completedAt *gorm.DeletedAt
	if completed {
		now := gorm.DeletedAt{Time: time.Now().UTC(), Valid: true}
		completedAt = &now
	}

	err := s.repo.UpdateStatusByCategory(category, completed, completedAt)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByCategory(category)
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
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

func (s *TodoService) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}

func (s *TodoService) DeleteAllTodos() error {
	return s.repo.DeleteAll()
}
