package user

import "github.com/gotripe/model"

type Repo interface {
	FindByID(UserID string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) (string, error)
	UpdateSubID(userID string, SubID string) error
}
