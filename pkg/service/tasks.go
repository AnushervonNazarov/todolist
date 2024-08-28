package service

import (
	"fmt"
	"todoList/logger"
	"todoList/models"
	"todoList/pkg/repository"
)

func GetAllTasks() ([]models.Task, error) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		logger.Error.Printf("[service.GetAllTasks] error getting all tasks: %v\n", err)
		return nil, fmt.Errorf("could not retrieve tasks: %v", err)
	}
	return tasks, nil
}

func GetTaskByID(id uint) (*models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		logger.Error.Printf("[service.GetTaskByID] error getting task: %v\n", err)
		return nil, fmt.Errorf("task not found: %v", err)
	}
	return task, nil
}

func AddTask(taskInput models.Task) (*models.Task, error) {
	newTask, err := repository.AddTask(&taskInput)
	if err != nil {
		logger.Error.Printf("[service.AddTask] error adding task: %v\n", err)
		return nil, fmt.Errorf("could not add task: %v", err)
	}
	return newTask, nil
}

func MarkTaskAsDoneByID(id uint) (*models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		logger.Error.Printf("[service.MarkTaskAsDoneByID] error marking task: %v\n", err)
		return nil, fmt.Errorf("task not found: %v", err)
	}

	task.IsDone = "true"

	updatedTask, err := repository.EditTaskByID(task)
	if err != nil {
		logger.Error.Printf("[service.MarkTaskAsDoneByID] error marking task: %v\n", err)
		return nil, fmt.Errorf("could not mark task as done: %v", err)
	}

	return updatedTask, nil
}

func EditTaskByID(id uint, taskInput models.Task) (*models.Task, error) {
	_, err := repository.GetTaskByID(id)
	if err != nil {
		logger.Error.Printf("[service.EditTaskByID] error editing task: %v\n", err)
		return nil, fmt.Errorf("task not found: %v", err)
	}

	updatedTask, err := repository.EditTaskByID(&taskInput)
	if err != nil {
		logger.Error.Printf("[service.EditTaskByID] error editing task: %v\n", err)
		return nil, fmt.Errorf("could not update task: %v", err)
	}

	taskInput.ID = id

	return updatedTask, nil
}

func DeleteTaskByID(id uint) error {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		logger.Error.Printf("[service.DeleteTaskByID] error deleating task: %v\n", err)
		return fmt.Errorf("task not found: %v", err)
	}

	if err := repository.DeleteTask(task); err != nil {
		logger.Error.Printf("[service.DeleteTaskByID] error deleating task: %v\n", err)
		return fmt.Errorf("could not delete task: %v", err)
	}

	return nil
}
