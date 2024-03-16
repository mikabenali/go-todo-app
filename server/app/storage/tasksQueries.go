package storage

import (
	"context"
	"errors"
	"main/types"

	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Database) GetTasksCollection() *mongo.Collection {
	return d.GetDatabase().Collection("tasks")
}

func (d *Database) GetAllTasks() ([]types.Task, error) {
	cursor, err := d.GetTasksCollection().Find(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	var tasks []types.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (d *Database) GetTaskById(id string) (types.Task, error) {
	for _, task := range d.Tasks {
		if task.Id == id {
			return task, nil
		}
	}

	return types.Task{}, errors.New("Task not found")
}

func (d *Database) CreateTask(task types.Task) error {
	d.Tasks = append(d.Tasks, task)

	return nil
}
