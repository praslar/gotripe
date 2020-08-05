package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gotripe/model"
	"github.com/gotripe/utils"
)

// SignUp register for a new account
func (h *Handler) CreateMovie(c echo.Context) error {
	var mv model.Movie
	req := &createMovie{}
	if err := req.bind(c, &mv); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.movieRepo.Create(&mv); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, &mv)
}

func (h *Handler) FindAllMovies(c echo.Context) error {
	articles, err := h.movieRepo.FindAll()
	if err != nil {
		c.Logger().Info(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *Handler) WatchMovie(c echo.Context) error {
	id := c.Param("movie_id")
	mv, err := h.movieRepo.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	c.Logger().Debug("Id: " + id + " | info: " + mv.Name)
	if mv == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, mv)
}
