package domain

import "errors"

const (
	maxTickerLength = 8
)

type Ticker struct {
	Symbol             string
	ShorName           string
	LongName           string
	Currency           string
	RegularMarketPrice float64
	EarningsPerShare   float64
}

func NewTicker(ticker string) (Ticker, error) {
	if ticker == "" {
		return Ticker{}, errors.New("invalid ticker")
	}

	if len(ticker) > maxTickerLength {
		return Ticker{}, errors.New("invalid ticker length")
	}

	return Ticker{
		Symbol: ticker,
	}, nil
}
