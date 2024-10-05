package main

import (
	"net/http"
	"net/mail"

	"github.com/krzysztofengineer/template/db"
	"github.com/krzysztofengineer/template/forms"
	"github.com/krzysztofengineer/template/pages"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	DB *db.Queries
}

func NewLoginHandler(db *db.Queries) *LoginHandler {
	return &LoginHandler{
		DB: db,
	}
}

func (*LoginHandler) Page(c echo.Context) error {
	return pages.Login().Render(c.Request().Context(), c.Response().Writer)
}

func (h *LoginHandler) Form(c echo.Context) error {
	email := c.FormValue("email")
	if email == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		errors := map[string]string{
			"email": "The email is required",
		}
		return forms.Login(errors).Render(c.Request().Context(), c.Response().Writer)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		errors := map[string]string{
			"email": "The email format is invalid",
		}
		return forms.Login(errors).Render(c.Request().Context(), c.Response().Writer)
	}

	h.DB.SaveUser(c.Request().Context(), email)

	return nil
}
