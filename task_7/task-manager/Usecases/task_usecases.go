package usecases

import (
	"errors"

	"github.com/kidistbezabih/task-manager/Domain"
	repositories "github.com/kidistbezabih/task-manager/Repositories"
)

type TaskUsecase struct {
	taskservice repositories.TaskServices
}

func NewTaskUsecase(taskservice repositories.TaskServices) TaskUsecase {
	return TaskUsecase{
		taskservice: taskservice,
	}
}

func (tu *TaskUsecase) GetAllTasks() ([]Domain.Task, error) {
	var tasks []Domain.Task
	tasks, err := tu.taskservice.GetAllTasks()

	if err != nil {
		return []Domain.Task{}, err
	}

	if len(tasks) == 0 {
		return []Domain.Task{}, errors.New("no task available")
	}
	return tasks, nil
}

func (tu *TaskUsecase) GetTaskById(id int) (Domain.Task, error) {
	task, err := tu.taskservice.GetTaskById(id)
	if err != nil {
		return Domain.Task{}, err
	}
	return task, err
}

func (tu *TaskUsecase) UpdateTaskById(task Domain.Task) error {
	err := tu.taskservice.UpdateTaskById(&task)

	return err
}

func (tu *TaskUsecase) DeleteTaskById(id int) error {
	err := tu.taskservice.DeleteTaskById(id)
	return err
}

func (tu *TaskUsecase) CreateTask(task *Domain.Task) error {
	err := tu.taskservice.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskUsecase) GetAUserTasks(username string) ([]Domain.Task, error) {
	tasks, err := tu.taskservice.GetAUserTasks(username)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
