package data

import (
	"context"
	"errors"
	"task_management/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskManagement interface {
	GetAllTasks() ([]model.Task, error)
	GetTaskById(id int) (model.Task, error)
	UpdateTaskById(task *model.Task) error
	DeleteTaskById(id int) error
	CreateTask(task *model.Task) error
}

type TaskServices struct {
	taskcollection *mongo.Collection
	ctx            context.Context
}

func NewTaskService(taskcollection *mongo.Collection, ctx context.Context) *TaskServices {
	return &TaskServices{
		taskcollection: taskcollection,
		ctx:            ctx,
	}
}

func (ts *TaskServices) GetAllTasks() ([]model.Task, error) {
	var allTasks []model.Task

	cursor, err := ts.taskcollection.Find(ts.ctx, bson.D{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ts.ctx)

	for cursor.Next(ts.ctx) {
		var task model.Task

		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		allTasks = append(allTasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(allTasks) == 0 {
		return nil, errors.New("no task")
	}
	return allTasks, nil
}

/// done

func (ts *TaskServices) GetTaskById(id int) (model.Task, error) {
	var task model.Task

	filter := bson.D{bson.E{Key: "id", Value: id}}
	err := ts.taskcollection.FindOne(ts.ctx, filter).Decode(&task)
	return task, err
}

func (ts *TaskServices) UpdateTaskById(task *model.Task) error {
	filter := bson.D{bson.E{Key: "id", Value: task.ID}}
	// update := bson.D{bson.E{Key: "$set", Value: bson.E{Key: "id", Value: updatedtask.ID}, bson.E{Key: "title", Value: updatedtask.Title}, bson.E{Key: "description", Value: updatedtask.Description}, bson.E{Key: "completed", Value: updatedtask.Completed}}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "id", Value: task.ID}, bson.E{Key: "title", Value: task.Title}, bson.E{Key: "description", Value: task.Description}, bson.E{Key: "completed", Value: task.Completed}}}}
	result, _ := ts.taskcollection.UpdateOne(ts.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no task with this id")
	}
	return nil
}

// func (ts *TaskServices)

func (ts *TaskServices) DeleteTaskById(id int) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, _ := ts.taskcollection.DeleteOne(ts.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no user with this id")
	}
	return nil
}

func (ts *TaskServices) CreateTask(task *model.Task) error {
	_, err := ts.taskcollection.InsertOne(ts.ctx, task)
	return err
}
