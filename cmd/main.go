package main

import (
	"log"

	"github.com/Alekcey5977/todo-app"
	"github.com/Alekcey5977/todo-app/pkg/handler"
	"github.com/Alekcey5977/todo-app/pkg/repository"
	"github.com/Alekcey5977/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	// Entry point for the todo application
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}