package handlers

import (
	"context"
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
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {

	taskRequest := request.Body

	taskToCreate := src.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.service.CreateTask(taskToCreate.Text, taskToCreate.IsDone)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) PatchTasksTaskId(ctx context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
	// Получаем задачу по ID
	taskToUpdate, err := h.service.GetTaskByID(*request.Body.Id)
	if err != nil {
		return nil, err
	}

	// Обновляем текст задачи и статус IsDone
	taskToUpdate.Text = *request.Body.Task
	taskToUpdate.IsDone = *request.Body.IsDone

	// Сохраняем обновленную задачу
	updatedTask, err := h.service.UpdateTask(taskToUpdate.ID, taskToUpdate.Text)
	if err != nil {
		return nil, err
	}

	// Возвращаем обновленную задачу в ответе
	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
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
