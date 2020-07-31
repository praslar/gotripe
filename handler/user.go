package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gotripe/db"
	"github.com/gotripe/model"
	"github.com/gotripe/utils"
)

func (h *Handler) SignUp(c echo.Context) error {
	var usr model.User
	req := &signUp{}
	if err := req.bind(c, &usr); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	dupUser, err := h.userRepo.FindByEmail(usr.Email)

	if err != nil && !db.IsErrNotFound(err) {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if dupUser != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.Duplicate())
	}

	_, err = h.userRepo.Create(&usr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&usr))
}
