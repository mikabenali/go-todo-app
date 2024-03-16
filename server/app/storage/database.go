package storage

import (
	"errors"
	"main/types"

	"github.com/google/uuid"
)

type Config struct{}

type Database struct {
	Config Config

	// For testing
	Tasks []types.Task
}

func NewDatabase(config Config) *Database {
	tasks := []types.Task{
		{Id: uuid.NewString(), Name: "My first task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My second task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My third task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My other task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My foo task", Description: "lorem ipsum lalala"},
		{Id: uuid.NewString(), Name: "My bar task", Description: "lorem ipsum lalala"},
	}

	return &Database{
		Config: config,
		Tasks:  tasks,
	}
}

func (d *Database) GetAllTasks() ([]types.Task, error) {
	return d.Tasks, nil
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
