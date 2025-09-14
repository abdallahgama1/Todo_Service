package models

import "database/sql"

type Todo struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Completed   bool           `json:"completed"`
	Category    string         `json:"category"`
	Priority    string         `json:"priority"`
	CompletedAt sql.NullString `json:"completedAt"`
	DueDate     sql.NullString `json:"dueDate"`
	UserID      uint            `json:"userId"`
	User        User            `json:"-"` 

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
