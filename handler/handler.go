package handler

import (
	"github.com/gotripe/movie"
	"github.com/gotripe/plan"
	"github.com/gotripe/sub"
	"github.com/gotripe/user"
)

type Handler struct {
	userRepo  user.Repo
	movieRepo movie.Repo
	subRepo   sub.Repo
	planRepo  plan.Repo
}

func NewHandler(usr user.Repo, mv movie.Repo, sub sub.Repo, plan plan.Repo) *Handler {
	return &Handler{
		movieRepo: mv,
		userRepo:  usr,
		planRepo:  plan,
		subRepo:   sub,
	}
}
