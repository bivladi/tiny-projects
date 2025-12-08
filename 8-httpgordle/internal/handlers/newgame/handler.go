package newgame

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	apiGame := api.GameResponse{}
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}
