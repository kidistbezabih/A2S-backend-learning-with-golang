package data

import (
	"context"
	"errors"
	"task_management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskManagement interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	UpdateTaskById(task *models.Task) error
	DeleteTaskById(id int) error
	CreateTask(task *models.Task, username string) error
	GetAUserTasks(username string) ([]models.Task, error)
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

func (ts *TaskServices) GetAllTasks() ([]models.Task, error) {
	var allTasks []models.Task

	cursor, err := ts.taskcollection.Find(ts.ctx, bson.D{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ts.ctx)

	for cursor.Next(ts.ctx) {
		var task models.Task

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

func (ts *TaskServices) GetTaskById(id int) (models.Task, error) {
	var task models.Task

	filter := bson.D{bson.E{Key: "id", Value: id}}
	err := ts.taskcollection.FindOne(ts.ctx, filter).Decode(&task)
	return task, err
}

func (ts *TaskServices) UpdateTaskById(task *models.Task) error {
	filter := bson.D{bson.E{Key: "id", Value: task.ID}}
	// update := bson.D{bson.E{Key: "$set", Value: bson.E{Key: "id", Value: updatedtask.ID}, bson.E{Key: "title", Value: updatedtask.Title}, bson.E{Key: "description", Value: updatedtask.Description}, bson.E{Key: "completed", Value: updatedtask.Completed}}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "id", Value: task.ID}, bson.E{Key: "title", Value: task.Title}, bson.E{Key: "description", Value: task.Description}, bson.E{Key: "completed", Value: task.Completed}}}}
	result, _ := ts.taskcollection.UpdateOne(ts.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no task with this id")
	}
	return nil
}

func (ts *TaskServices) DeleteTaskById(id int) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, _ := ts.taskcollection.DeleteOne(ts.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no user with this id")
	}
	return nil
}

func (ts *TaskServices) CreateTask(task *models.Task, username string) error {
	task.Username = username
	_, err := ts.taskcollection.InsertOne(ts.ctx, task)
	return err
}

func (us *TaskServices) GetAUserTasks(username string) ([]models.Task, error) {
	filter := bson.D{bson.E{Key: "username", Value: username}}
	cursor, err := us.taskcollection.Find(us.ctx, filter)
	var tasks []models.Task

	if err != nil {
		return nil, err
	}
	defer cursor.Close(us.ctx)

	for cursor.Next(us.ctx) {
		var task models.Task

		err := cursor.Decode(&task)
		if err != nil {
			return []models.Task{}, err
		}
		tasks = append(tasks, task)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("no user with the username")
	}

	return tasks, nil
}
