package dtos

import "moneybits/core/modules/stocks/domain"

type TickerDataResponse struct {
	Symbol             string  `json:"symbol,omitempty"`
	ShorName           string  `json:"shor_name,omitempty"`
	LongName           string  `json:"long_name,omitempty"`
	Currency           string  `json:"currency,omitempty"`
	RegularMarketPrice float64 `json:"regular_market_price"`
	EarningsPerShare   float64 `json:"earnings_per_share"`
}

func NewTickerDataResponse(source domain.Ticker) TickerDataResponse {
	return TickerDataResponse{
		Symbol:             source.Symbol,
		ShorName:           source.ShorName,
		LongName:           source.LongName,
		Currency:           source.Currency,
		RegularMarketPrice: source.RegularMarketPrice,
		EarningsPerShare:   source.EarningsPerShare,
	}
}
