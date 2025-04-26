package domain

import "errors"

const (
	maxTickerLength = 8
)

type Ticker struct {
	Code string
}

func NewTicker(ticker string) (Ticker, error) {
	if ticker == "" {
		return Ticker{}, errors.New("invalid ticker")
	}

	if len(ticker) > maxTickerLength {
		return Ticker{}, errors.New("invalid ticker length")
	}

	return Ticker{
		Code: ticker,
	}, nil
}
