package types

type Task struct {
	Id          string `json:"_id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type TaskRequest struct {
	Name        string `json:"name" bson:"name" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
}
