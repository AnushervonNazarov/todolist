package repository

import (
	"todoList/db"
	"todoList/logger"
	"todoList/models"
)

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := db.GetDBConn().Find(&tasks).Error; err != nil {
		logger.Error.Printf("[repository.GetAllTasks] error getting all tasks: %v\n", err)
		return nil, translateError(err)
	}
	return tasks, nil
}

func GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := db.GetDBConn().First(&task, id).Error; err != nil {
		logger.Error.Printf("[repository.GetAllTaskByID] error getting task by id: %v\n", err)
		return nil, translateError(err)
	}
	return &task, nil
}

func AddTask(task *models.Task) (*models.Task, error) {
	if err := db.GetDBConn().Create(task).Error; err != nil {
		logger.Error.Printf("[repository.AddTask] error adding task: %v\n", err)
		return nil, translateError(err)
	}
	return task, nil
}

func EditTaskByID(task *models.Task) (*models.Task, error) {
	if err := db.GetDBConn().Save(task).Error; err != nil {
		logger.Error.Printf("[repository.EditTaskByID] error editing task: %v\n", err)
		return nil, translateError(err)
	}
	return task, nil
}

func DeleteTask(task *models.Task) error {
	if err := db.GetDBConn().Delete(task).Error; err != nil {
		logger.Error.Printf("[repository.DeleteTask] error deleating task: %v\n", err)
		return translateError(err)
	}
	return nil
}
