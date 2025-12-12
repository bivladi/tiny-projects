package handlers

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/getstatus"
	"learngo/httpgordle/internal/handlers/guess"
	"learngo/httpgordle/internal/handlers/newgame"
	"learngo/httpgordle/internal/repository"
	"net/http"
)

// NewRouter returns a router that listens for requests
// to the following endpoints:
//   - Create a new game;
//
// The provided router is ready to serve.
func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db))
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db))
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handler(db))
	return r
}
