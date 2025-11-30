package gordle

import (
	"testing"
)

func TestFeedbackString(t *testing.T) {
	tt := map[string]struct{
		input feedback
		want string
	} {
		"empty": {
			input: feedback{},
			want: "",
		},
		"one hint": {
			input: feedback{0},
			want: "⬜️",
		},
		"different hints": {
			input: feedback{0,1,2},
			want: "⬜️🟡💚",
		},
		"wrong hint": {
			input: feedback{4},
			want: "💔",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			if tc.input.String() != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, tc.input.String())
			}
		})
	}
}