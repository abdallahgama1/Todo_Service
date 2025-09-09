// repositories/todo_repository.go

package repositories

import (
	"Todo_Service/models"
	"strings"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo

	result := r.db.First(&todo, id)
	return &todo, result.Error
}

func (r *TodoRepository) GetByCategory(category string) ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Where("category = ?", category).Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) GetByStatus(completed bool) ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Where("completed = ?", completed).Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) SearchByTitle(query string) ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(query)+"%").Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) Create(todo *models.Todo) (*models.Todo, error) {

	result := r.db.Create(todo)
	return todo, result.Error
}

func (r *TodoRepository) Update(todo *models.Todo) (*models.Todo, error) {

	result := r.db.Save(todo)
	return todo, result.Error
}

func (r *TodoRepository) UpdateStatusByCategory(category string, completed bool, completedAt *gorm.DeletedAt) error {

	result := r.db.Model(&models.Todo{}).Where("category = ?", category).Updates(map[string]interface{}{"completed": completed, "completed_at": completedAt})
	return result.Error
}

func (r *TodoRepository) Delete(id uint) error {

	result := r.db.Delete(&models.Todo{}, id)
	return result.Error
}

func (r *TodoRepository) DeleteAll() error {
	result := r.db.Unscoped().Where("1 = 1").Delete(&models.Todo{})
	return result.Error
}
