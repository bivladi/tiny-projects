package money_test

import (
	"learngo/moneyconverter/money"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount money.Amount
		to money.Currency
		validate func (t *testing.T, got money.Amount, err error)
	} {
		"39.98 EUR to USD": {
			amount: money.Amount{},
			to: money.Currency{},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Fatalf("expected no error, got %s", err.Error())
				}
				expected := money.Amount{}
				if !reflect.DeepEqual(got, expected) {
					t.Fatalf("expected %v, got %v", expected, got)
				}
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			amount,err := money.Convert(tc.amount, tc.to)
			tc.validate(t, amount, err)
		})
	}
}
