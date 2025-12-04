package money

import (
	"errors"
	"testing"
)

func TestParseCurrency(t *testing.T) {
	tt := map[string]struct {
		input    string
		expected Currency
		err      error
	}{
		"euro": {
			input:    "EUR",
			expected: Currency{code: "EUR", precision: 2},
			err:      nil,
		},
		"IRR test": {
			input:    "IRR",
			expected: Currency{code: "IRR", precision: 0},
			err:      nil,
		},
		"VND test": {
			input:    "VND",
			expected: Currency{code: "VND", precision: 1},
			err:      nil,
		},
		"OMR test": {
			input:    "OMR",
			expected: Currency{code: "OMR", precision: 3},
			err:      nil,
		},
		"Unknown currency": {
			input:    "value",
			expected: Currency{},
			err:      ErrInvalidCurrencyCode,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.input)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}
			// ✔ we can compare because struct contains comparable types
			if got != tc.expected {
				t.Fatalf("expected currency %v, got %v", tc.expected.code, got.code)
			}
		})
	}
}
