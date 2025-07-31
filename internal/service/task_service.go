package service

import (
	"errors"
	"strings"
	"time"
	"togo-cli/internal/domain"
	"togo-cli/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateNewTask(description, priority string) (*domain.Task, error) {
	description = strings.TrimSpace(description)
	if description == "" {
		return nil, errors.New("Description can't be empty")
	}

	priority = strings.TrimSpace(priority)
	if priority == "" {
		priority = "normal"
	}

	newTask := &domain.Task{
		Id:          -1,
		Priority:    priority,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	err := s.repo.Create(newTask)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (s *TaskService) Complete(task *domain.Task) error {
	if task.Completed {
		return errors.New("Task is already completed")
	}

	task.Completed = true
	s.repo.Update(task)
	return nil
}

func (s *TaskService) GetTasks() ([]*domain.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetTasksByStatus(status string) ([]*domain.Task, error) {
	return s.repo.GetAllByStatus(status)
}

func (s *TaskService) GetTaskByID(id int) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Update(task *domain.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) Delete(id int) error {
	return s.repo.Delete(id)
}
