package handler

import (
	"github.com/gotripe/model"
	"github.com/gotripe/utils"
)

type userResponse struct {
	User struct {
		Email    string `json:"email"`
		PLan     string `json:"plan"`
		StripeID string `json:"cus_id"`
		Token    string `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Email = u.Email
	r.User.PLan = u.Plan
	r.User.StripeID = u.StripeID
	r.User.Token = utils.GenerateJWT(u.UserID)
	return r
}
