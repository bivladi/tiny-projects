package money

type Amount struct {
	amount Decimal
	currency Currency
}

const (
    // ErrTooPrecise is returned if the number is 
    // too precise for the currency.
    ErrTooPrecise = Error("quantity is too precise")
)

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}
	quantity.precision = currency.precision
	return Amount{amount: quantity, currency: currency}, nil
}