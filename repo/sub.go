package repo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gotripe/model"
)

type (
	SubRepo struct {
		session *mgo.Session
	}
)

func NewSubRepo(session *mgo.Session) *SubRepo {
	return &SubRepo{
		session: session,
	}
}

func (r *SubRepo) FindByID(SubID string) (*model.Subscription, error) {
	selector := bson.M{"sub_id": SubID}
	s := r.session.Clone()
	defer s.Close()
	var sub *model.Subscription
	if err := r.collection(s).Find(selector).One(&sub); err != nil {
		return nil, err
	}
	return sub, nil
}

func (r *SubRepo) Create(sub *model.Subscription) error {
	s := r.session.Clone()
	defer s.Clone()

	return r.collection(s).Insert(sub)

}
func (r *SubRepo) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("sub")
}
