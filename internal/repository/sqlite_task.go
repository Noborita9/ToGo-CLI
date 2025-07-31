package repository

import (
	"database/sql"
	"fmt"
	"slices"
	"togo-cli/internal/domain"
)

type SQLiteTaskRepository struct {
	db *sql.DB
}

func NewSQLiteTaskRepository(db *sql.DB) *SQLiteTaskRepository {
	return &SQLiteTaskRepository{db: db}
}

func (s *SQLiteTaskRepository) Create(task *domain.Task) error {
	query := `
		INSERT INTO tasks(completed, priority, description, created_at)
		VALUES (?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		task.Completed,
		task.Priority,
		task.Description,
		task.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("Could not create task in databse: %w", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("Could not get id from new task: %w", err)
	}

	task.Id = int(id)

	return nil
}

func (s *SQLiteTaskRepository) GetAll() ([]*domain.Task, error) {
	query := `
		SELECT id, completed, priority, description, created_at FROM tasks ORDER BY created_at ASC;
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Could not get tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]*domain.Task, 0)
	for rows.Next() {
		task := &domain.Task{}
		err := rows.Scan(
			&task.Id,
			&task.Completed,
			&task.Priority,
			&task.Description,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Could not parse tasks: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *SQLiteTaskRepository) GetAllByStatus(status string) ([]*domain.Task, error) {

	validStatus := []string{"pending", "completed"}

	isValid := slices.Contains(validStatus, status)

	if !isValid {
		return nil, fmt.Errorf("Status is not valid please use pending|completed.")
	}

	query := `
		SELECT id, completed, priority, description, created_at FROM tasks WHERE completed=? ORDER BY created_at ASC;
	`

	rows, err := s.db.Query(
		query,
		status == "completed",
	)
	if err != nil {
		return nil, fmt.Errorf("Could not get tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]*domain.Task, 0)
	for rows.Next() {
		task := &domain.Task{}
		err := rows.Scan(
			&task.Id,
			&task.Completed,
			&task.Priority,
			&task.Description,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Could not parse tasks: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *SQLiteTaskRepository) GetByID(id int) (*domain.Task, error) {
	query := `
	SELECT id, completed, priority, description, created_at FROM tasks WHERE id=?;
	`
	task := &domain.Task{}
	err := s.db.QueryRow(query, id).Scan(
		&task.Id,
		&task.Completed,
		&task.Priority,
		&task.Description,
		&task.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Task with id: %d not found", id)
		}
		return nil, fmt.Errorf("Connection failed to get task: %w", err)
	}

	return task, nil
}

func (s *SQLiteTaskRepository) Update(task *domain.Task) error {
	query := `
		UPDATE tasks SET completed=?, priority=?, description=? WHERE id=?
	`
	_, err := s.db.Exec(query,
		task.Completed,
		task.Priority,
		task.Description,
		task.Id,
	)

	if err != nil {
		return fmt.Errorf("Could not update task: %w", err)
	}

	return nil
}

func (s *SQLiteTaskRepository) Delete(id int) error {
	query := `
		DELETE FROM tasks WHERE id=?
	`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Could not delete task: %w", err)
	}
	return nil
}
