package money

import (
	"errors"
	"testing"
)

func TestNewAmount(t *testing.T) {
	euroCurrency := Currency{code: "EUR", precision: 2}
	tt := map[string]struct {
		quantity Decimal
		currency Currency
		want Amount
		err error 
	} {
		"10.50 Euro": {
			quantity: Decimal{subunits: 1050, precision: 2},
			currency: euroCurrency,
			want: Amount{amount: Decimal{subunits: 1050, precision: 2}, currency: euroCurrency},
			err: nil,
		},
		"10.001 Euro": {
			quantity: Decimal{subunits: 10001, precision: 3},
			currency: euroCurrency,
			want: Amount{},
			err: ErrTooPrecise,
		},
	}
	for name, tc := range tt {
		t.Run(name, func (t *testing.T) {
			got, err := NewAmount(tc.quantity, tc.currency)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.want {
				t.Fatalf("expected amount %v, got %v", tc.want, got)
			}
		})
	}
}