package repositories

import (
	"Todo_Service/models"
	"database/sql"
	"fmt"
	"strings"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) InitDB() error {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        completed BOOLEAN NOT NULL,
        category VARCHAR(50) NOT NULL,
        priority VARCHAR(10) NOT NULL,
        completedAt TIMESTAMPTZ,
        dueDate TIMESTAMPTZ
    );`
	_, err := r.db.Exec(query)
	return err
}

func scanTodos(rows *sql.Rows) ([]models.Todo, error) {
	todos := []models.Todo{}
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.Category, &t.Priority, &t.CompletedAt, &t.DueDate); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todos ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTodos(rows)
}

func (r *TodoRepository) GetByID(id int) (*models.Todo, error) {
	var t models.Todo
	err := r.db.QueryRow("SELECT * FROM todos WHERE id = $1", id).Scan(&t.ID, &t.Title, &t.Completed, &t.Category, &t.Priority, &t.CompletedAt, &t.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, err
	}
	return &t, nil
}

func (r *TodoRepository) GetByCategory(category string) ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todos WHERE category = $1 ORDER BY id ASC", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTodos(rows)
}

func (r *TodoRepository) GetByStatus(completed bool) ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todos WHERE completed = $1 ORDER BY id ASC", completed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTodos(rows)
}

func (r *TodoRepository) SearchByTitle(query string) ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todos WHERE LOWER(title) LIKE $1 ORDER BY id ASC", "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTodos(rows)
}

func (r *TodoRepository) Create(todo models.Todo) (*models.Todo, error) {
	query := `
    INSERT INTO todos (title, completed, category, priority, dueDate, completedAt)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id`

	var id int
	err := r.db.QueryRow(query, todo.Title, todo.Completed, todo.Category, todo.Priority, todo.DueDate, todo.CompletedAt).Scan(&id)
	if err != nil {
		return nil, err
	}
	todo.ID = id
	return &todo, nil
}

func (r *TodoRepository) Update(id int, todo models.Todo) (*models.Todo, error) {
	query := `
    UPDATE todos
    SET title = $1, completed = $2, category = $3, priority = $4, dueDate = $5, completedAt = $6
    WHERE id = $7`

	result, err := r.db.Exec(query, todo.Title, todo.Completed, todo.Category, todo.Priority, todo.DueDate, todo.CompletedAt, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("todo not found")
	}

	todo.ID = id
	return &todo, nil
}

func (r *TodoRepository) UpdateStatusByCategory(category string, completed bool, completedAt sql.NullString) error {
	query := `UPDATE todos SET completed = $1, completedAt = $2 WHERE category = $3`
	_, err := r.db.Exec(query, completed, completedAt, category)
	return err
}

func (r *TodoRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}

func (r *TodoRepository) DeleteAll() error {
	_, err := r.db.Exec("DELETE FROM todos")
	return err
}
