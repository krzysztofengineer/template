package main

import (
	"github.com/krzysztofengineer/template/pages"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return pages.Home().Render(c.Request().Context(), c.Response().Writer)

	})
	e.Logger.Fatal(e.Start(":3000"))
}
