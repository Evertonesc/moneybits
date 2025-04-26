package core

import "github.com/labstack/echo/v4"

// HTTPRouter provides the application HTTP router to all
// components that need it.
func (app *AppContainer) HTTPRouter() *echo.Router {
	return app.HTTPServer.Router()
}
