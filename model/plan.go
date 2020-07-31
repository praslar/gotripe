package model

type (
	Plan struct {
		PlanID        string `json:"plan_id,omitempty" bson:"movie_id,omitempty"`
		Name          string `json:"name,omitempty" bson:"name,omitempty"`
		Active        bool   `json:"active,omitempty" bson:"active,omitempty"`
		Amount        int    `json:"amount,omitempty" bson:"amount,omitempty"`
		Currency      string `json:"currency,omitempty" bson:"currency,omitempty"`
		Interval      string `json:"interval,omitempty" bson:"interval"`
		IntervalCount int    `json:"interval_count,omitempty" bson:"interval_count"`
	}
)
