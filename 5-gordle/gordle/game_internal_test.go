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
			g,_ := New(strings.NewReader(tt.input), []string{string(tt.want)}, 0)
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
			g,_ := New(nil, []string{"SLICE"}, 0)
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

func TestGame_computeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess []rune
		solution []rune
		expectedFeedback feedback
	} {
		"no matches": {
			guess: []rune("AAAAA"),
			solution: []rune("BBBBB"),
			expectedFeedback: feedback{
				absentCharacter,
				absentCharacter,
				absentCharacter,
				absentCharacter,
				absentCharacter,
			},
		},
		"one match": {
			guess: []rune("AAAAA"),
			solution: []rune("ABBBB"),
			expectedFeedback: feedback{
				correctPosition,
				absentCharacter,
				absentCharacter,
				absentCharacter,
				absentCharacter,
			},
		},
		"partial match": {
			guess: []rune("AAAAC"),
			solution: []rune("ABBCB"),
			expectedFeedback: feedback{
				correctPosition,
				absentCharacter,
				absentCharacter,
				absentCharacter,
				wrongPosition,
			},
		},
		"all match": {
			guess: []rune("AAAAA"),
			solution: []rune("AAAAA"),
			expectedFeedback: feedback{
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		"same letter in different places": {
			guess: []rune("hello"),
			solution: []rune("hlleo"),
			expectedFeedback: feedback{
				correctPosition,
				wrongPosition,
				correctPosition,
				wrongPosition,
				correctPosition,
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			result := computeFeedback(tc.guess, tc.solution)
			if !result.Equal(tc.expectedFeedback) {
				t.Fatalf("expected %v, got %v", tc.expectedFeedback, result)
			}
		})
	}
}