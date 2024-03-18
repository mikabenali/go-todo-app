package storage

import (
	"context"
	"main/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Database) GetTasksCollection() *mongo.Collection {
	return d.GetDatabase().Collection("tasks")
}

func (d *Database) GetAllTasks() ([]types.Task, error) {
	cursor, err := d.GetTasksCollection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var tasks []types.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (d *Database) GetTaskById(id primitive.ObjectID) (types.Task, error) {
	var task types.Task

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	if err := d.GetTasksCollection().FindOne(context.TODO(), filter).Decode(&task); err != nil {
		return types.Task{}, err
	}

	return task, nil
}

func (d *Database) CreateTask(task *types.TaskRequest) (types.Task, error) {
	result, err := d.GetTasksCollection().InsertOne(context.TODO(), task)
	if err != nil {
		return types.Task{}, err
	}

	return d.GetTaskById(result.InsertedID.(primitive.ObjectID))
}

func (d *Database) UpdateTask(task types.TaskRequest, id primitive.ObjectID) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := d.GetTasksCollection().UpdateOne(context.TODO(), filter, task)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteTask(id primitive.ObjectID) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := d.GetTasksCollection().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
