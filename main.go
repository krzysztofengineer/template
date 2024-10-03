package main

import (
	"embed"

	"github.com/krzysztofengineer/template/pages"
	"github.com/labstack/echo/v4"
)

var (
	//go:embed dist
	distFS embed.FS
)

func main() {
	e := echo.New()

	e.StaticFS("dist", echo.MustSubFS(distFS, "dist"))

	e.GET("/", func(c echo.Context) error {
		return pages.Home().Render(c.Request().Context(), c.Response().Writer)

	})

	e.Logger.Fatal(e.Start(":3000"))
}
