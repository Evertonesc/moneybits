package brapi

import (
	"context"
	"fmt"
	"moneybits/drivers/envs"
	"moneybits/drivers/rest"
)

type RestAdapter interface {
	Get(ctx context.Context, url string, headers, queryParams map[string]string, response any) error
}

type BrapiAdapter struct {
	bi RestAdapter
}

func NewBrapiAdapter() *BrapiAdapter {
	return &BrapiAdapter{
		bi: rest.NewRestClient(envs.EnvConfig.BrapiURL),
	}
}

func (ba *BrapiAdapter) FetchTickerQuote(ctx context.Context, ticker string) (TickerQuoteResponse, error) {
	url := fmt.Sprintf("/quote/%s", ticker)

	headers := map[string]string{
		rest.AuthorizationHeader: envs.EnvConfig.BrapiToken,
	}

	var tickerQuoteResponse TickerQuoteResponse
	err := ba.bi.Get(ctx, url, headers, nil, &tickerQuoteResponse)
	if err != nil {
		return TickerQuoteResponse{}, err
	}

	return tickerQuoteResponse, nil
}
