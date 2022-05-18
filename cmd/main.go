package main

import (
	"github.com/zinchenkom/go-todo-app"
	"github.com/zinchenkom/go-todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server %s", err.Error())
	}
}
