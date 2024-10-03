package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
}

func (*DashboardHandler) Index(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/login")
}
