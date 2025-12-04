package money

type Amount struct {
	amount   Decimal
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is
	// too precise for the currency.
	ErrTooPrecise = Error("quantity is too precise")
)

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	switch {
	case quantity.precision > currency.precision:
		// In order to avoid converting 0.00001 cent, let's exit now.
		return Amount{}, ErrTooPrecise
	case quantity.precision < currency.precision:
		quantity.subunits *= pow10(currency.precision - quantity.precision)
		quantity.precision = currency.precision
	}
	return Amount{amount: quantity, currency: currency}, nil
}

func (a Amount) String() string {
	return a.amount.String() + " " + a.currency.code
}

func (a Amount) validate() error {
	switch {
	case a.amount.subunits > maxDecimal:
		return ErrTooLarge
	case a.amount.precision > a.currency.precision:
		return ErrTooPrecise
	}
	return nil
}
