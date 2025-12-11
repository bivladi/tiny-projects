package main

import (
	"learngo/httpgordle/internal/handlers"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8082", handlers.NewRouter())
	if err != nil {
		panic(err)
	}
}
