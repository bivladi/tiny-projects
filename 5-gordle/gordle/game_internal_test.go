package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGame_ask(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in English": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tt.input))
			got := g.ask()
			if !slices.Equal(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestGame_validateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("POCKET"),
			expected: errInvalidWordLength,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil)
			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf(
					"%c, expected %q, got %q",
					tc.word, tc.expected, err,
				)
			}
		})
	}
}
