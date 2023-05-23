package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/efrenfuentes/todo-backend-go/internal/validator"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Url       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TodoModel struct {
	DB *sql.DB
}

func ValidateTodo(v *validator.Validator, todo *Todo) {
	v.Check(todo.Title != "", "title", "must be provided")
	v.Check(len(todo.Title) <= 100, "title", "must not be more than 100 bytes long")
}

// The Insert() method will insert a new Todo record into the database.
func (m *TodoModel) Insert(todo *Todo) error {
	query := `
		INSERT INTO todos (title, completed, "order", url)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	args := []interface{}{todo.Title, todo.Completed, todo.Order, todo.Url}

	return m.DB.QueryRow(query, args...).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
}

func (m *TodoModel) Get(id int64) (*Todo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, title, completed, "order", url, created_at, updated_at
		FROM todos
		WHERE id = $1
	`

	var todo Todo

	err := m.DB.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.Order,
		&todo.Url,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &todo, nil
}

func (m *TodoModel) GetAll() ([]*Todo, error) {
	query := `
		SELECT id, title, completed, "order", url, created_at, updated_at
		FROM todos
		ORDER BY id`

	// Create a context with a 3-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() to execute the query with the context we created above.
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	// Importantly, we defer a call to rows.Close() to ensure that the sql.Rows result set is
	// always properly closed before the GetAll() method returns.
	defer rows.Close()

	todos := []*Todo{}

	for rows.Next() {
		// Initialize a new Todo instance for each row.
		var todo Todo

		// Scan the row into the Todo struct.
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.Order,
			&todo.Url,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Append the todo to the slice of todos.
		todos = append(todos, &todo)
	}

	// Check for errors from the rows.Next() iterator.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went well, return the slice of todos.
	return todos, nil
}

func (m *TodoModel) Update(todo *Todo) error {
	query := `
		UPDATE todos
		SET title = $1, completed = $2, "order" = $3, url = $4, updated_at = $5
		WHERE id = $6
		RETURNING updated_at`

	args := []interface{}{todo.Title, todo.Completed, todo.Order, todo.Url, time.Now(), todo.ID}

	return m.DB.QueryRow(query, args...).Scan(&todo.UpdatedAt)
}

func (m *TodoModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM todos
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, then the record must not exist.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
