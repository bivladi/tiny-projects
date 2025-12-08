package handlers

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/newgame"
	"net/http"
)

// NewRouter returns a router that listens for requests
// to the following endpoints:
//   - Create a new game;
//
// The provided router is ready to serve.
func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(
		http.MethodPost+" "+api.NewGameRoute,
		newgame.Handle,
	)
	return r
}

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}
