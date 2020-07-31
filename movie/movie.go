package movie

import "github.com/gotripe/model"

type Repo interface {
	FindAll() ([]*model.Movie, error)
	Create(movie *model.Movie) error
}
