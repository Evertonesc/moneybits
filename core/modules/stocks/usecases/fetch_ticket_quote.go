package usecases

import (
	"context"
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

func (uc *FetchTickerQuoteUC) Execute(ctx context.Context, ticker domain.Ticker) (*brapi.TickerQuoteResponse, error) {
	return uc.ba.FetchTickerQuote(ctx, ticker)
}
