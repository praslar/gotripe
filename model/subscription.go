package model

import (
	"time"
)

type (
	Subscription struct {
		SubID  string    `json:"sub_id,omitempty" bson:"sub_id,omitempty"`
		PlanID string    `json:"plan_id,omitempty" bson:"movie_id,omitempty"`
		Start  time.Time `json:"start" bson:"start"`
		End    time.Time `json:"end" bson:"end"`
	}
)
