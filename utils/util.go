package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func WithUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("user", "mukund")
		// ctx := context.WithValue(c.Request().Context(), "user", "pro@mukund.com")
		// c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
