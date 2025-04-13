package api

import (
	"context"
	"moneybits/adapters/brapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FetchTickerQuoteUseCase interface {
	Execute(ctx context.Context, ticker string) (brapi.TickerQuoteResponse, error)
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
	ticker := c.Param("ticker")
	reqCtx := c.Request().Context()

	resp, err := h.uc.Execute(reqCtx, ticker)
	if err != nil {
		// create a global error handler to automatic return the right status and body
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
