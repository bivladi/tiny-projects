package ecbank

import (
	"fmt"
	"learngo/moneyconverter/money"
	"net/http"
)

type Client struct {
	url string
}

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrClientSide         = ecbankError("client side error")
	ErrServerSide         = ecbankError("server side error")
	ErrUnknownStatusCode  = ecbankError("unknown status code")
	ErrUnexpectedFormat   = ecbankError("unexpected format")
	ErrChangeRateNotFound = ecbankError("change rate is not found")
	path                  = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	clientErrorClass      = 4
	serverErrorClass      = 5
)

func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	if c.url == "" {
		c.url = path
	}
	resp, err := http.Get(c.url)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w:%s", ErrCallingServer, err.Error())
	}
	defer resp.Body.Close()
	if err := checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}
	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}
	return rate, nil
}

func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:

		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w:%d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w:%d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w:%d", ErrUnknownStatusCode, statusCode)
	}
}

func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}
