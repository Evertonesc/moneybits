package usecases

import (
	"context"
	"errors"
	"moneybits/adapters/brapi"
	"moneybits/core/modules/stocks/domain"
)

type BrapiAdapter interface {
	FetchTickerQuote(ctx context.Context, ticker domain.Ticker) (*brapi.TickerQuoteResponse, error)
}

type FetchTickerQuoteUC struct {
	ba BrapiAdapter
}

func NewFetchTickerQuoteUC(ba BrapiAdapter) *FetchTickerQuoteUC {
	return &FetchTickerQuoteUC{
		ba: ba,
	}
}

func (uc *FetchTickerQuoteUC) Execute(ctx context.Context, ticker domain.Ticker) (domain.Ticker, error) {
	ticketData, err := uc.ba.FetchTickerQuote(ctx, ticker)
	if err != nil {
		// TODO: create apperrs package to handle errors
		return domain.Ticker{}, err
	}

	if len(ticketData.Results) > 1 {
		return domain.Ticker{}, errors.New("unprocessable data")
	}

	result := ticketData.Results[0]
	return domain.Ticker{
		Symbol:             result.Symbol,
		ShorName:           result.ShortName,
		LongName:           result.LongName,
		Currency:           result.Currency,
		RegularMarketPrice: result.RegularMarketPrice,
		EarningsPerShare:   result.EarningsPerShare,
	}, nil
}
