package handlers

import (
	"context"
	"fmt"

	// "fmt"
	"goapi/internal/src"
	"goapi/internal/web/tasks"
	// "log"
)

type TaskHandler struct {
	service src.TaskService
}

func NewTaskHandler(s src.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	// Проверка на nil
	if taskRequest.Text == nil || *taskRequest.Text == "" {
		return nil, fmt.Errorf("invalid task: text cannot be empty")
	}
	if taskRequest.IsDone == nil {
		return nil, fmt.Errorf("invalid task: is_done cannot be nil")
	}

	// Создаем задачу
	createdTask, err := h.service.CreateTask(*taskRequest.Text, *taskRequest.IsDone)
	if err != nil {
		return nil, err
	}

	// Возвращаем ответ
	return tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Text:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}, nil
}

func (h *TaskHandler) PatchTasksTaskId(ctx context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
	// Получаем задачу по ID
	taskToUpdate, err := h.service.GetTaskByID(*request.Body.Id)
	if err != nil {
		return nil, err
	}

	// Обновляем текст задачи и статус IsDone
	taskToUpdate.Text = *request.Body.Text
	taskToUpdate.IsDone = *request.Body.IsDone

	// Сохраняем обновленную задачу
	updatedTask, err := h.service.UpdateTask(taskToUpdate.ID, taskToUpdate.Text, taskToUpdate.IsDone)
	if err != nil {
		return nil, err
	}

	// Возвращаем обновленную задачу в ответе
	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &updatedTask.ID,
		Text:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksTaskId(ctx context.Context, request tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	taskId := request.TaskId

	err := h.service.DeleteTask(uint(taskId))
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksTaskId204Response{}, nil
}
