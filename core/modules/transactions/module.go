package transactions

import (
	"moneybits/core/modules/transactions/api"

	"github.com/labstack/echo/v4"
)

type AppContainer interface {
	HTTPRouter() *echo.Router
}

type TransactionModule struct {
	app AppContainer
}

func NewTransactionModule(app AppContainer) *TransactionModule {
	return &TransactionModule{
		app: app,
	}
}

func (m *TransactionModule) Name() string {
	return "transactions"
}

func (m *TransactionModule) Start() error {
	createTXHandler := api.NewCreateTransaction(nil)

	api.RegisterTransactionRoutes(m.app.HTTPRouter(), createTXHandler)

	return nil
}
