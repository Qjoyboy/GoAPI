package handlers

import (
	"GoApi/internal/src"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service src.TaskService
}

func NewTaskHandler(s src.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTask(c echo.Context) error {
	var req src.TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task, err := h.service.CreateTask(req.Text, false)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) PatchTask(c echo.Context) error {
	var req src.TaskRequest
	id := c.Param("id")
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "not found task"})
	}

	updatedTask, err := h.service.UpdateTask(id, req.Text)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}
