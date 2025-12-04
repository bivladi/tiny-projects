package main

import (
	"flag"
	"fmt"
	"learngo/moneyconverter/money"
	"os"
)

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")
	flag.Parse()
	fromCurrency := parseCurrency(from)
	toCurrency := parseCurrency(to)
	value := parseDecimal(flag.Arg(0))
	amount, err := money.NewAmount(value, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Amount:", amount, "; Currency: ", toCurrency)
	convertedAmount, err := money.Convert(amount, toCurrency)
	if err != nil {
		_, _ = fmt.Fprintf(
			os.Stderr,
			"unable to convert %s to %s: %s\n",
			amount, toCurrency, err.Error(),
		)
		os.Exit(1)
	}
	fmt.Printf("%s = %s\n", amount, convertedAmount)
}

func parseCurrency(input *string) money.Currency {
	currency, err := money.ParseCurrency(*input)
	if err != nil {
		_, _ = fmt.Fprintf(
			os.Stderr,
			"unable to parse currency %q: %s\n",
			*input, err.Error(),
		)
		os.Exit(1)
	}
	return currency
}

func parseDecimal(input string) money.Decimal {
	decimal, err := money.ParseDecimal(input)
	if err != nil {
		_, _ = fmt.Fprintf(
			os.Stderr,
			"unable to parse decimal %q: %s\n",
			input, err.Error(),
		)
	}
	return decimal
}
