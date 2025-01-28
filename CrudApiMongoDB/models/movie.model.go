package movieModel

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Movie struct {
	ID      bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string        `json:"movie,omitempty"`
	Watched bool          `json:"watched,omitempty"`
}
