package main

import (
	"goapi/internal/db"
	"goapi/internal/handlers"
	"goapi/internal/src"
	usersrc "goapi/internal/userSrc"
	"goapi/internal/web/tasks"
	"goapi/internal/web/users"
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
	taskService := src.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := usersrc.NewTaskRepository(database)
	userService := usersrc.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	UHandler := users.NewStrictHandler(userHandler, nil)
	strictHander := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHander)
	users.RegisterHandlers(e, UHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
