package rest

import (
	"moneybits/drivers/envs"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func NewHTTPServer() *echo.Echo {
	echoSrv := echo.New()

	echoSrv.Use(
		middleware.CORSWithConfig(corsSettings()),
		middleware.TimeoutWithConfig(timeoutSettings()),
		middleware.RateLimiterWithConfig(rateLimitSettings()),
		middleware.RequestID(),
	)

	return echoSrv
}

func corsSettings() middleware.CORSConfig {
	return middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		// TODO: change the origins persmission after the first version
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		MaxAge:       envs.EnvConfig.CorsMaxAge,
	}
}

func timeoutSettings() middleware.TimeoutConfig {
	return middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request timeout: The server took too long to respond. Please try again in a few minutes",
		Timeout:      envs.EnvConfig.HTTPServerTimeout,
	}
}

func rateLimitSettings() middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10), Burst: 20, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
}
