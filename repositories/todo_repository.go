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

func (r *TodoRepository) GetAllForUser(userId uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where("user_id = ?", userId).Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByIDForUser(id, userId uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.Where("id = ? AND user_id = ?", id, userId).First(&todo).Error
	return &todo, err
}

func (r *TodoRepository) GetByCategoryForUser(category string, userId uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where("category = ? AND user_id = ?", category, userId).Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByStatusForUser(completed bool, userId uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where("completed = ? AND user_id = ?", completed, userId).Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) SearchByTitleForUser(query string, userId uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where("user_id = ? AND LOWER(title) LIKE ?", userId, "%"+strings.ToLower(query)+"%").Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	err := r.db.Create(todo).Error
	return todo, err
}

func (r *TodoRepository) Update(todo *models.Todo) (*models.Todo, error) {
	err := r.db.Save(todo).Error
	return todo, err
}

func (r *TodoRepository) Delete(todo *models.Todo) error {
	return r.db.Delete(todo).Error
}

func (r *TodoRepository) UpdateStatusByCategoryForUser(category string, completed bool, completedAt *gorm.DeletedAt, userId uint) error {
	return r.db.Model(&models.Todo{}).
		Where("category = ? AND user_id = ?", category, userId).
		Updates(map[string]interface{}{"completed": completed, "completed_at": completedAt}).Error
}

func (r *TodoRepository) DeleteAllForUser(userId uint) error {
	return r.db.Unscoped().Where("user_id = ?", userId).Delete(&models.Todo{}).Error
}
