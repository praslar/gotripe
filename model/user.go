package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		UserID       string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
		SubID        string    `json:"sub_id,omitempty" bson:"sub_id"`
		Email        string    `json:"email,omitempty" bson:"email,omitempty"`
		Plan         string    `json:"plan,omitempty" bson:"plan,omitempty"`
		StripeID     string    `json:"stripe_customer_id,omitempty" bson:"stripe_customer_id,omitempty"`
		SessionToken string    `json:"session_token" bson:"session_token"`
		SubStatus    string    `json:"sub_status,omitempty" bson:"sub_status,omitempty"`
		Password     string    `json:"-" bson:"password,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdateAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}
)

func (u *User) HashPassword(plain string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
