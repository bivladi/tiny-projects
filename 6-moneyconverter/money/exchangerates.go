package money

type ratesFetcher interface {
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}
