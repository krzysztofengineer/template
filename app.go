package main

import (
	"embed"

	"github.com/krzysztofengineer/template/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

var (
	//go:embed dist
	distFS embed.FS
)

func NewApp() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(30))))

	e.StaticFS("dist", echo.MustSubFS(distFS, "dist"))

	e.GET("/", func(c echo.Context) error {
		return pages.Home().Render(c.Request().Context(), c.Response().Writer)
	})

	return e
}
