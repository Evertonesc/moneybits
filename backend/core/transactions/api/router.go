package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterTransactionRoutes(router *echo.Router, createTXHandler *CreateTransactionHandler) {
	router.Add(
		http.MethodPost,
		"/transactions",
		createTXHandler.Create,
	)
}
