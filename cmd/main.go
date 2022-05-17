package main

import (
	"github.com/zinchenkom/go-todo-app"
	"log"
)


func main(){
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error occured while running http server %s", err.Error())
	}
}
