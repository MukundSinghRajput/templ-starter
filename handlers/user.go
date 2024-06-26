package handlers

import (
	"templ-starter/model"
	"templ-starter/utils"
	"templ-starter/views/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (u UserHandler) HandlerUserShow(c echo.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		us := model.User{
			Name:   "Mukund",
			Email:  "pro@mukund.com", // ofc not real mail id
			UserID: "7019511932",
		}
		return utils.Render(c, user.Show(us, c))
	}
}
