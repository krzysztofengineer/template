package main

import (
	"database/sql"

	"github.com/krzysztofengineer/template/pages"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	DB *sql.DB
}

func NewLoginHandler(db *sql.DB) *LoginHandler {
	return &LoginHandler{
		DB: db,
	}
}

func (*LoginHandler) Page(c echo.Context) error {
	return pages.Login().Render(c.Request().Context(), c.Response().Writer)
}
