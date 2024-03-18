package api

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetObjectIdFromString(id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return objectId, nil
}
