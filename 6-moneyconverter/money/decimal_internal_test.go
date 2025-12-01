package money

import (
	"errors"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tt := map[string]struct{
		decimal string
		expected Decimal
		err error
	} {
		"2 decimal digits": {
            decimal:  "1.52",
            expected: Decimal{subunits: 152, precision: 2},
            err:      nil,
        },
        "no decimal digits": {
			decimal: "12",
			expected: Decimal{subunits: 12, precision: 0},
			err: nil,
		},
        "suffix 0 as decimal digits": {
			decimal: "10.0",
			expected: Decimal{subunits: 100, precision: 1},
			err: nil,
		},
        "prefix 0 as decimal digits": {
			decimal: "10.01",
			expected: Decimal{subunits: 1001, precision: 2},
			err: nil,
		},
        "multiple of 10": {
			decimal: "101010",
			expected: Decimal{subunits: 101010, precision: 0},
			err: nil,
		},
        "invalid decimal part": {
			decimal: "12.",
			expected: Decimal{subunits: 12, precision: 0},
			err: nil,
		},
        "not a number": {
            decimal: "NaN",
            err:    ErrInvalidDecimal,
        },
        "empty string": {
            decimal: "",
            err:    ErrInvalidDecimal,
        },
        "too large": {
            decimal: "1234567890123",
            err:    ErrTooLarge,
        },
	}
	for name, tc := range tt {
		t.Run(name, func (t *testing.T) {
			got,err := ParseDecimal(tc.decimal)
			if !errors.Is(tc.err, err) {
				t.Fatalf("expected: %v, got: %v", tc.err, err)
			}
			if tc.expected.subunits != got.subunits {
				t.Fatalf("expected: %d, got: %d", tc.expected.subunits, got.subunits)
			}
			if tc.expected.precision != got.precision {
				t.Fatalf("expected: %d, got: %d", tc.expected.subunits, got.subunits)
			}
		})
	}
}