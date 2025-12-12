package newgame

import (
	"learngo/httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/games", nil)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()

	handleFunc := Handle(gameAdderStub{})
	handleFunc.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(
		t,
		`{"attempts_left":0, "guesses":[], "id":"", "status":"", "word_length":0}`,
		recorder.Body.String(),
	)
}

type gameAdderStub struct {
	err error
}

func (g gameAdderStub) Add(_ session.Game) error {
	return g.err
}
