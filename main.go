package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	Text string `json:"text"`
}

type TaskRequest struct {
	Text string `json:"text"`
}

var tasks = []Task{}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task := Task{
		Text: req.Text,
	}

	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, fmt.Sprintf("Hello, %s", task.Text))
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTask)
	e.POST("/tasks", postTask)

	e.Start("localhost:8080")
}
