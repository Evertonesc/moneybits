package core

import (
	"fmt"
	"log"
	"moneybits/core/modules/stocks"
	"moneybits/core/modules/transactions"
	"moneybits/drivers/envs"
	"moneybits/drivers/rest"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Module represents a self-contained feature of the application.
// A module provides APIs and use cases to the outer layer and also manages its own dependencies
type Module interface {
	// Name returns the module identifier
	Name() string
	// Start initializes the module (background tasks, etc)
	Start() error
}

// AppContainer is responsible for managing the application's core dependencies and lifecycle.
// It handles:
//   - Database connections and their lifecycle
//   - Application startup and graceful shutdown
//   - Environment configuration management
//   - Service container and dependency resolution
type AppContainer struct {
	HTTPServer *echo.Echo
	Modules    []Module
	EnvConfig  envs.Config
}

func NewAppContainer() *AppContainer {
	envs.Envs()
	echoServer := rest.NewHTTPServer()

	app := &AppContainer{
		HTTPServer: echoServer,
		Modules:    make([]Module, 0),
		EnvConfig:  envs.EnvConfig,
	}

	app.bootstrap()

	return app
}

func (app *AppContainer) bootstrap() {
	app.modulesRegistry()

	for _, module := range app.Modules {
		err := module.Start()
		if err != nil {
			log.Fatalf("error initilizing module %s: %s", module.Name(), err.Error())
		}
	}
}

func (app *AppContainer) modulesRegistry() {
	modulesList := []Module{
		transactions.NewTransactionModule(app),
		stocks.NewSocksModule(app),
	}

	app.Modules = modulesList
}

func (app *AppContainer) StartHTTPServer() {
	err := app.HTTPServer.Start(fmt.Sprintf(":%s", app.EnvConfig.HTTPServerPort))
	if err == http.ErrServerClosed {
		log.Fatalf("http server closed: %s", err.Error())
	}
}

// HTTPRouter provides the application HTTP router to all
// components that need it.
func (app *AppContainer) HTTPRouter() *echo.Router {
	return app.HTTPServer.Router()
}
