package newgame

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
)

type gameAdder interface {
	Add(game session.Game) error
}

func Handle(db gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		game, err := createGame(db)
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(w, "failed to create a new game", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		apiGame := api.ToGameResponse(game)
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func createGame(db gameAdder) (session.Game, error) {
	game := session.Game{}
	err := db.Add(game)
	if err != nil {
		return session.Game{}, err
	}
	return game, nil
}
