package api

import (
	"context"
	"moneybits/core/modules/stocks/api/dtos"
	"moneybits/core/modules/stocks/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FetchTickerQuoteUseCase interface {
	Execute(ctx context.Context, ticker domain.Ticker) (domain.Ticker, error)
}

type FetchTickerQuoteHandler struct {
	uc FetchTickerQuoteUseCase
	// presenter here
}

func NewFetchTickerQuoteHandler(uc FetchTickerQuoteUseCase) *FetchTickerQuoteHandler {
	return &FetchTickerQuoteHandler{
		uc: uc,
	}
}

func (h *FetchTickerQuoteHandler) FetchTickerQuote(c echo.Context) error {
	ctx := c.Request().Context()
	ticker := c.Param("ticker")

	newTicker, err := domain.NewTicker(ticker)
	if err != nil {
		return err
	}

	tickerData, err := h.uc.Execute(ctx, newTicker)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dtos.NewTickerDataResponse(tickerData))
}
