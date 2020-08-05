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
		Plan     string `json:"plan"`
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
	usr.Plan = req.User.Plan
	if err != nil {
		return err
	}
	usr.Password = hashpwd
	return nil
}

type Login struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *Login) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type createMovie struct {
	Movie struct {
		Name        string `json:"name" validate:"required"`
		Amount      int    `json:"amount" validate:"required"`
		ThumbNail   string `json:"thumb,omitempty" `
		Description string `json:"description"`
	} `json:"movie"`
}

func (r *createMovie) bind(c echo.Context, mv *model.Movie) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	mv.Name = r.Movie.Name
	mv.Amount = r.Movie.Amount
	mv.Description = r.Movie.Description
	mv.MovieID = utils.NewID()
	mv.ThumbNail = r.Movie.ThumbNail
	return nil
}
