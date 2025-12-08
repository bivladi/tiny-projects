package newgame

import (
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

	Handle(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(
		t,
		`{"attempts_left":0, "guesses":null, "id":"", "status":"", "word_length":0}`,
		recorder.Body.String(),
	)
}
