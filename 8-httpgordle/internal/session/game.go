package session

import (
	"errors"
	"learngo/httpgordle/internal/gordle"
)

type Game struct {
	ID           GameID
	Gordle       gordle.Game
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

type GameID string

type Status string

type Guess struct {
	Word     string
	Feedback string
}

const (
	StatusPlaying = "Playing"
	StatusWon     = "Won"
	StatusLost    = "Lost"
)

var ErrGameOver = errors.New("game over")
