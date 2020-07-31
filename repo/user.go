package repo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gotripe/model"
)

type (
	UserRepo struct {
		session *mgo.Session
	}
)

func NewUserRepo(session *mgo.Session) *UserRepo {
	return &UserRepo{
		session: session,
	}
}

// FindByID find user by uuid
func (r *UserRepo) FindByID(UserID string) (*model.User, error) {
	selector := bson.M{"user_id": UserID}
	s := r.session.Clone()
	defer s.Close()
	var users *model.User
	if err := r.collection(s).Find(selector).One(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// FindByEmail find user by email to prevent duplicate user
func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	selector := bson.M{"email": email}
	s := r.session.Clone()
	defer s.Close()
	var user *model.User
	if err := r.collection(s).Find(selector).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// Create new user in database
func (r *UserRepo) Create(user *model.User) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	user.CreatedAt = time.Now()
	user.UpdateAt = user.CreatedAt

	if err := r.collection(s).Insert(user); err != nil {
		return "", err
	}
	return user.UserID, nil
}

// UpdateSubID or create new SubID for user
func (r *UserRepo) UpdateSubID(userID string, SubID string) error {
	s := r.session.Clone()
	defer s.Clone()
	return r.collection(s).Update(bson.M{"user_id": userID}, bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
		"$addToSet": bson.M{
			"sub_id": SubID,
		},
	},
	)
}

func (r *UserRepo) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("user")
}
