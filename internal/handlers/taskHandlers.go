package handlers

import (
	"context"
	"fmt"
	"goapi/internal/src"
	"goapi/internal/web/tasks"
)

type TaskHandler struct {
	service src.TaskService
}

// GetUsersUserIdTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) GetUsersUserIdTasks(ctx context.Context, request tasks.GetUsersUserIdTasksRequestObject) (tasks.GetUsersUserIdTasksResponseObject, error) {
	userId := request.UserId
	userTasks, err := h.service.GetTasksForUser(userId)
	if err != nil {
		return nil, err
	}
	response := tasks.GetUsersUserIdTasks200JSONResponse{}

	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func NewTaskHandler(s src.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// PostTasksUserId implements tasks.StrictServerInterface.
func (h *TaskHandler) PostTasksUserId(ctx context.Context, request tasks.PostTasksUserIdRequestObject) (tasks.PostTasksUserIdResponseObject, error) {
	taskRequest := request.Body

	createdTask, err := h.service.CreateTaskByUserId(*taskRequest.Text, *taskRequest.UserId, *taskRequest.IsDone)
	if err != nil {
		return nil, err
	}

	return tasks.PostTasksUserId201JSONResponse{
		Id:     &createdTask.ID,
		Text:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}, nil

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

	if taskRequest.Text == nil || *taskRequest.Text == "" {
		return nil, fmt.Errorf("invalid task: text cannot be empty")
	}
	if taskRequest.IsDone == nil {
		return nil, fmt.Errorf("invalid task: is_done cannot be nil")
	}

	createdTask, err := h.service.CreateTask(*taskRequest.Text, *taskRequest.IsDone)
	if err != nil {
		return nil, err
	}

	return tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Text:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}, nil
}

func (h *TaskHandler) PatchTasksTaskId(ctx context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {

	taskId := request.TaskId
	taskToUpdate, err := h.service.GetTaskByID(taskId)
	if err != nil {
		return nil, err
	}

	taskToUpdate.Text = *request.Body.Text
	taskToUpdate.IsDone = *request.Body.IsDone

	updatedTask, err := h.service.UpdateTask(taskToUpdate.ID, taskToUpdate.Text, taskToUpdate.IsDone)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &updatedTask.ID,
		Text:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksTaskId(ctx context.Context, request tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	taskId := request.TaskId

	err := h.service.DeleteTask(string(taskId))
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksTaskId204Response{}, nil
}
