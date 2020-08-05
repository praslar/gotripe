package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/customer"

	"github.com/gotripe/db"
	"github.com/gotripe/model"
	"github.com/gotripe/utils"
)

// SignUp register for a new account
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

	// create stripe customer
	var conf Config
	utils.Load(&conf)

	stripe.Key = conf.SecKey
	params := &stripe.CustomerParams{
		Email: stripe.String(usr.Email),
	}

	cus, err := customer.New(params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	usr.StripeID = cus.ID

	_, err = h.userRepo.Create(&usr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&usr))
}

// Login for existing user
func (h *Handler) Login(c echo.Context) error {
	req := &Login{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.userRepo.FindByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}
