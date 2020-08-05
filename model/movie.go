package model

import "time"

type (
	Movie struct {
		MovieID     string    `json:"movie_id,omitempty" bson:"movie_id,omitempty"`
		Name        string    `json:"name,omitempty" bson:"name,omitempty"`
		Amount      int       `json:"amount,omitempty" bson:"amount,omitempty"`
		Description string    `json:"description,omitempty" bson:"description,omitempty"`
		ThumbNail   string    `json:"thumb,omitempty" bson:"thumb,omitempty"`
		CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdateAt    time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	}
)
