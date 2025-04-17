package main

import (
	"GoApi/internal/db"
	"GoApi/internal/handlers"
	"GoApi/internal/src"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := src.NewTaskRepository(database)
	taskService := src.NewTaskSerivce(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandler.GetTask)
	e.POST("/tasks", taskHandler.PostTask)
	e.PATCH("/tasks/:id", taskHandler.PatchTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	e.Start("localhost:8080")
}
