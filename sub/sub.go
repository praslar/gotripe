package sub

import "github.com/gotripe/model"

type Repo interface {
	FindByID(SubID string) (*model.Subscription, error)
	Create(sub *model.Subscription) error
}
