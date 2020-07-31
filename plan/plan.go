package plan

import "github.com/gotripe/model"

type Repo interface {
	FindByID(PlanID string) (*model.Plan, error)
	FindAll() ([]*model.Plan, error)
	Create(plan *model.Plan) error
}
