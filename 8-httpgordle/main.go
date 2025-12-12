package main

import (
	"learngo/httpgordle/internal/handlers"
	"learngo/httpgordle/internal/repository"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8082", handlers.NewRouter(repository.New()))
	if err != nil {
		panic(err)
	}
}
