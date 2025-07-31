package repository

import "togo-cli/internal/domain"

type TaskRepository interface {
	Create(task *domain.Task) error
	GetAll() ([]*domain.Task, error)
	GetAllByStatus(status string) ([]*domain.Task, error)
	GetByID(id int) (*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id int) error
}
