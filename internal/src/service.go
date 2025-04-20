package src

import (
	"fmt"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(text string, is_done bool) (Task, error)
	GetTasks() ([]Task, error)
	GetTaskByID(id uint) (Task, error)
	UpdateTask(id uint, text string) (Task, error)
	DeleteTask(id uint) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) taskValidate(text string, is_done bool) (string, bool) {
	if text == "" {
		return "", false
	}
	return text, is_done
}

func (s *taskService) CreateTask(text string, is_done bool) (Task, error) {
	text, isValid := s.taskValidate(text, is_done)
	if !isValid {
		return Task{}, fmt.Errorf("invalid task: text cannot be empty")
	}

	task := Task{
		ID:   uint(uuid.New().ID()),
		Text: text,
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
func (s *taskService) UpdateTask(id uint, text string) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	isdone := false
	text, is_done := s.taskValidate(text, isdone)

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
