package repo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gotripe/model"
)

type (
	MovieRepo struct {
		session *mgo.Session
	}
)

func NewMovieRepo(session *mgo.Session) *MovieRepo {
	return &MovieRepo{
		session: session,
	}
}

func (r *MovieRepo) FindAll() ([]*model.Movie, error) {
	s := r.session.Clone()
	defer s.Close()
	var movies []*model.Movie
	if err := r.collection(s).Find(nil).All(&movies); err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepo) FindByID(MovieID string) (*model.Movie, error) {
	selector := bson.M{"movie_id": MovieID}
	s := r.session.Clone()
	defer s.Close()
	var movie *model.Movie
	if err := r.collection(s).Find(selector).One(&movie); err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *MovieRepo) Create(movie *model.Movie) error {
	s := r.session.Clone()
	defer s.Clone()

	movie.CreatedAt = time.Now()
	movie.UpdateAt = movie.CreatedAt

	return r.collection(s).Insert(movie)

}
func (r *MovieRepo) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("movie")
}
