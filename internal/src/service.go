package src

import (
	"fmt"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(text string, is_done bool) (Task, error)
	GetTasks() ([]Task, error)
	GetTaskByID(id uint) (Task, error)
	UpdateTask(id uint, text string, is_done bool) (Task, error)
	DeleteTask(id uint) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) taskValidate(text string) error {
	if text == "" {
		return fmt.Errorf("text cannot be empty")
	}
	return nil
}

func (s *taskService) CreateTask(text string, is_done bool) (Task, error) {
	if err := s.taskValidate(text); err != nil {
		return Task{}, err
	}

	task := Task{
		ID:     uint(uuid.New().ID()),
		Text:   text,
		IsDone: is_done,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

// GetAllTasks implements TaskService.
func (s *taskService) GetTasks() ([]Task, error) {
	return s.repo.GetTasks()
}

// GetTaskByID implements TaskService.
func (s *taskService) GetTaskByID(id uint) (Task, error) {
	return s.repo.GetTaskByID(id)
}

// UpdateTask implements TaskService.
func (s *taskService) UpdateTask(id uint, text string, is_done bool) (Task, error) {
	if err := s.taskValidate(text); err != nil {
		return Task{}, err
	}

	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	task.Text = text
	task.IsDone = is_done

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

// DeleteTask implements TaskService.
func (s *taskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
