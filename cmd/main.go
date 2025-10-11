package main

import (
	"log"

	"github.com/Alekcey5977/todo-app"
	"github.com/Alekcey5977/todo-app/pkg/handler"
)

func main() {
	// Entry point for the todo application
	handler := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run(":8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}