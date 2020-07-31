package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/gotripe/model"
	"github.com/gotripe/utils"
)

type signUp struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gt=6"`
	} `json:"user"`
}

func (req *signUp) bind(c echo.Context, usr *model.User) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	usr.UserID = utils.NewID()
	usr.Email = req.User.Email
	hashpwd, err := usr.HashPassword(req.User.Password)
	if err != nil {
		return err
	}
	usr.Password = hashpwd
	return nil
}
