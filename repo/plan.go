package repo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gotripe/model"
)

type (
	PlanRepo struct {
		session *mgo.Session
	}
)

func NewPlanRepo(session *mgo.Session) *PlanRepo {
	return &PlanRepo{
		session: session,
	}
}

func (r *PlanRepo) FindAll() ([]*model.Plan, error) {
	s := r.session.Clone()
	defer s.Close()
	var plans []*model.Plan
	if err := r.collection(s).Find(nil).All(&plans); err != nil {
		return nil, err
	}
	return plans, nil
}

func (r *PlanRepo) FindByID(PlanID string) (*model.Plan, error) {
	selector := bson.M{"Plan_id": PlanID}
	s := r.session.Clone()
	defer s.Close()
	var plans *model.Plan
	if err := r.collection(s).Find(selector).One(&plans); err != nil {
		return nil, err
	}
	return plans, nil
}

func (r *PlanRepo) Create(plan *model.Plan) error {
	s := r.session.Clone()
	defer s.Clone()

	return r.collection(s).Insert(plan)

}
func (r *PlanRepo) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("plan")
}
