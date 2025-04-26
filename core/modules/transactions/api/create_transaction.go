package api

import (
	"context"
	"moneybits/core/modules/transactions/dtos"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateTXUseCase interface {
	Execute(ctx context.Context, createTransactionReq dtos.CreateTransactionRequest) (dtos.CreateTransactionResponse, error)
}

type CreateTransactionHandler struct {
	uc CreateTXUseCase
}

func NewCreateTransaction(uc CreateTXUseCase) *CreateTransactionHandler {
	return &CreateTransactionHandler{
		uc: uc,
	}
}

func (h *CreateTransactionHandler) Create(c echo.Context) error {
	var request dtos.CreateTransactionRequest
	err := c.Bind(&request)
	if err != nil {
		return c.String(http.StatusBadRequest, "parse error: invalid body request")
	}

	return c.JSON(http.StatusCreated, nil)
}
