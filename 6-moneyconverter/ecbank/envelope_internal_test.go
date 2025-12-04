package ecbank

import (
	"errors"
	"learngo/moneyconverter/money"
	"reflect"
	"testing"
)

func TestExchangeRates(t *testing.T) {
	tt := map[string]struct {
		env  envelope
		want map[string]float64
	}{
		"nominal": {
			env: envelope{Rates: []currencyRate{
				{Rate: 2, Currency: "USD"},
				{Rate: 3, Currency: "CHF"},
				{Rate: 4, Currency: "AED"},
			}},
			want: map[string]float64{
				"USD": 2,
				"CHF": 3,
				"AED": 4,
				"EUR": 1,
			},
		},
		"empty": {
			env: envelope{[]currencyRate{}},
			want: map[string]float64{
				"EUR": 1,
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.env.exchangeRates()
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestExchangeRate(t *testing.T) {
	e := envelope{Rates: []currencyRate{
		{Rate: 2.2, Currency: "USD"},
		{Rate: 8.8, Currency: "AED"},
	}}
	tt := map[string]struct {
		env    envelope
		source string
		target string
		want   money.ExchangeRate
		err    error
	}{
		"nominal": {
			env:    e,
			source: baseCurrencyCode,
			target: "USD",
			want:   mustParceExchangeRate(t, "2.2000000000"),
		},
		"same currency": {
			env:    e,
			source: "USD",
			target: "USD",
			want:   mustParceExchangeRate(t, "1"),
		},
		"USD to AED": {
			env:    e,
			source: "USD",
			target: "AED",
			want:   mustParceExchangeRate(t, "4.0000000000"),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			rate, err := tc.env.exchangeRate(tc.source, tc.target)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %v, got %v", tc.err, err)
			}
			if rate != tc.want {
				t.Fatalf("expected %v, got %v", tc.want, rate)
			}
		})
	}
}

func mustParceExchangeRate(t *testing.T, rate string) money.ExchangeRate {
	t.Helper()

	decimal, err := money.ParseDecimal(rate)
	if err != nil {
		t.Fatalf("unable to parse decimal %s, err: %v", rate, err)
	}
	return money.ExchangeRate(decimal)
}
