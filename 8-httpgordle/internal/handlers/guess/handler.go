package guess

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
)

type gameGuesser interface {
	Find(game session.GameID) (session.Game, error)
	Update(game session.Game) error
}

func Handle(guesser gameGuesser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusBadRequest)
			return
		}
		log.Printf("guess next word for game with id: %v", id)
		request := api.GuessRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response := api.ToGameResponse(game(id, request, guesser))
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func game(id string, r api.GuessRequest, db gameGuesser) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
