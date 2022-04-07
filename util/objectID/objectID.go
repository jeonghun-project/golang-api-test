package objectID

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Hex(hex string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(hex)
	return id
}
