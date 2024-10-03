package main

import (
	"database/sql"
	"embed"

	"github.com/krzysztofengineer/template/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pressly/goose/v3"
	"golang.org/x/time/rate"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed dist
	distFS embed.FS

	//go:embed db/migrations/*.sql
	migrationsFS embed.FS
)

func main() {
	db, err := sql.Open("sqlite3", "./db/db.sqlite")
	if err != nil {
		panic(err)
	}

	e := NewApp(db)

	e.Logger.Fatal(e.Start(":3000"))
}

func NewApp(db *sql.DB) *echo.Echo {
	goose.SetBaseFS(migrationsFS)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "db/migrations"); err != nil {
		panic(err)
	}

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
