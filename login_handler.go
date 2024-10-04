package main

import (
	"database/sql"
	"net/http"

	"github.com/krzysztofengineer/template/forms"
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

func (*LoginHandler) Form(c echo.Context) error {
	c.Response().WriteHeader(http.StatusBadRequest)
	errors := map[string]string{
		"email": "The email is required",
	}
	return forms.Login(errors).Render(c.Request().Context(), c.Response().Writer)
}
