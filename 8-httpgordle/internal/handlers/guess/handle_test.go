package guess

import (
	"learngo/httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"learngo/httpgordle/internal/api"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodPost,
		"/games/",
		strings.NewReader(`{"guess": "word"}`),
	)
	require.NoError(t, err)
	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()
	handleFunc := Handle(gameGuesserStub{})
	handleFunc.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}

type gameGuesserStub struct {
	error
}

func (g gameGuesserStub) Find(game session.GameID) (session.Game, error) {
	return session.Game{}, g.error
}

func (g gameGuesserStub) Update(game session.Game) error {
	return nil
}
