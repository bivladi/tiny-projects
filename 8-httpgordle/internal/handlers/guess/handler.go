package guess

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
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
	response := api.ToGameResponse(game(id, request))
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}

func game(id string, r api.GuessRequest) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
