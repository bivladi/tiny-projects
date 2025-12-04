package money_test

import (
	"learngo/moneyconverter/money"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: mustParseAmount(t, "34.98", "USD"),
			to:     mustParseCurrency(t, "EUR"),
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Fatalf("expected no error, got %s", err.Error())
				}
				expected := mustParseAmount(t, "69.96", "EUR")
				if !reflect.DeepEqual(got, expected) {
					t.Fatalf("expected %v, got %v", expected, got)
				}
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			amount, err := money.Convert(tc.amount, tc.to)
			tc.validate(t, amount, err)
		})
	}
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("unable to parse currency code: %v", code)
	}
	return currency
}

func mustParseAmount(t *testing.T, value string, code string) money.Amount {
	decimal, err := money.ParseDecimal(value)
	if err != nil {
		t.Fatalf("unable to parse decimal %v", value)
	}
	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("unable to parse currency %v", code)
	}
	amount, err := money.NewAmount(decimal, currency)
	if err != nil {
		t.Fatalf("unable to create amount %s", err.Error())
	}
	return amount
}
