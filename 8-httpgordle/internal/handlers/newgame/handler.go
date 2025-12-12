package newgame

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/gordle"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"

	"github.com/oklog/ulid/v2"
)

type gameAdder interface {
	Add(game session.Game) error
}

func Handler(db gameAdder) http.HandlerFunc {
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

const maxAttempts = 5

func createGame(db gameAdder) (session.Game, error) {
	corpus, err := gordle.ParseCorpus()
	if err != nil {
		return session.Game{}, err
	}
	if len(corpus) == 0 {
		return session.Game{}, gordle.ErrEmptyCorpus
	}
	solution, err := gordle.PickRandomWord(corpus)
	if err != nil {
		return session.Game{}, err
	}
	game, err := gordle.New(solution)
	if err != nil {
		return session.Game{}, err
	}
	g := session.Game{
		ID:           session.GameID(ulid.Make().String()),
		Gordle:       *game,
		AttemptsLeft: maxAttempts,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
	}
	err = db.Add(g)
	if err != nil {
		return session.Game{}, err
	}
	return g, nil
}
