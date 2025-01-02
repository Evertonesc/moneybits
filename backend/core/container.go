package core

import (
	"fmt"
	"log"
	"moneybits/core/transactions"
	"moneybits/drivers/rest"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Module represents a self-contained feature of the application
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
}

func NewAppContainer() *AppContainer {
	echoServer := rest.NewHTTPServer()

	app := &AppContainer{
		HTTPServer: echoServer,
		Modules:    make([]Module, 0),
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
	}

	app.Modules = modulesList
}

func (app *AppContainer) StartHTTPServer() {
	err := app.HTTPServer.Start(fmt.Sprintf(":%s", "8080"))
	if err == http.ErrServerClosed {
		log.Fatalf("http server closed: %s", err.Error())
	}
}

func (app *AppContainer) HTTPRouter() *echo.Router {
	return app.HTTPServer.Router()
}