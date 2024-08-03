package data

import (
	// "Task_management/model"
	"errors"
	"example/task-management/model"
)

type TaskManagement interface {
	GetTasks() []model.Task
	GetTaskById(id int) (model.Task, error)
	UpdateTaskById(id int, task model.Task) error
	DeleteTaskById(id int) error
	CreateTask(task model.Task) model.Task
}

type TaskServices struct {
	nextId int
	tasks  map[int]model.Task
}

func NewTaskService() *TaskServices {
	return &TaskServices{
		nextId: 1,
		tasks:  make(map[int]model.Task),
	}
}

func (ts *TaskServices) GetTasks() []model.Task {
	var allTasks []model.Task

	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}

	return allTasks
}

func (ts *TaskServices) GetTaskById(id int) (model.Task, error) {
	taskById, exist := ts.tasks[id]

	if !exist {
		return model.Task{}, errors.New("there no task with this id")
	}
	return taskById, nil
}

func (ts *TaskServices) UpdateTaskById(id int, task model.Task) error {
	taskById, exist := ts.tasks[id]

	if !exist {
		return errors.New("there no task with this id")
	}
	if taskById.Description != task.Description {
		taskById.Description = task.Description
	}
	if taskById.Title != task.Title {
		taskById.Title = task.Title
	}
	if taskById.Completed != task.Completed {
		taskById.Completed = task.Completed
	}

	ts.tasks[id] = taskById
	return nil
}

// func (ts *TaskServices)

func (ts *TaskServices) DeleteTaskById(id int) error {
	taskById, exist := ts.tasks[id]

	if !exist {
		return errors.New("there no task with this id")
	}
	delete(ts.tasks, taskById.ID)
	return nil
}

func (ts *TaskServices) CreateTask(task model.Task) model.Task {

	task.ID = ts.nextId
	ts.tasks[task.ID] = task
	ts.nextId++
	return task
}
