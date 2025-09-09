package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model        // Embedded struct
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	Category   string `json:"category"`
	Priority   string `json:"priority"`

	CompletedAt *gorm.DeletedAt `json:"completedAt" gorm:"index"`
	DueDate     *gorm.DeletedAt `json:"dueDate"`
}

type CreateTodoRequest struct {
	Title     string  `json:"title" binding:"required"`
	Completed bool    `json:"completed"`
	Category  string  `json:"category" binding:"required"`
	Priority  string  `json:"priority" binding:"required"`
	DueDate   *string `json:"dueDate"`
}

type UpdateTodoRequest struct {
	Title     string  `json:"title" binding:"required"`
	Completed bool    `json:"completed"`
	Category  string  `json:"category" binding:"required"`
	Priority  string  `json:"priority" binding:"required"`
	DueDate   *string `json:"dueDate"`
}

type UpdateCategoryStatusRequest struct {
	Completed bool `json:"completed"`
}
