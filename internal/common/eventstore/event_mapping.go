package eventstore

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Mapping[T any](event any) (T, error) {
	var t T
	err := bson.Unmarshal(event.(primitive.Binary).Data, &t)
	if err != nil {
		return t, err
	}
	return t, err
}
