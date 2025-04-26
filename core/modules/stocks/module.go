package stocks

import (
	"moneybits/adapters/brapi"
	"moneybits/core/modules/stocks/api"
	"moneybits/core/modules/stocks/usecases"

	"github.com/labstack/echo/v4"
)

type AppContainer interface {
	HTTPRouter() *echo.Router
}

type StocksModule struct {
	app AppContainer
}

func NewSocksModule(app AppContainer) *StocksModule {
	return &StocksModule{
		app: app,
	}
}
func (m *StocksModule) Name() string {
	return "stocks"
}
func (m *StocksModule) Start() error {
	brapiAdapter := brapi.NewBrapiAdapter()
	fetchTickerQuoteUC := usecases.NewFetchTickerQuoteUC(brapiAdapter)
	fetchTickerQuoteHandler := api.NewFetchTickerQuoteHandler(fetchTickerQuoteUC)

	api.RegisterStocksRouter(m.app.HTTPRouter(), fetchTickerQuoteHandler)

	return nil
}
