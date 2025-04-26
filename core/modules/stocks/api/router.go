package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterStocksRouter(router *echo.Router, fetchTickerDataHandler *FetchTickerQuoteHandler) {
	router.Add(
		http.MethodGet,
		"/ticker/:ticker",
		fetchTickerDataHandler.FetchTickerQuote,
	)
}
