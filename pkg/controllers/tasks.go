package controllers

import (
	"net/http"
	"strconv"
	"todoList/logger"
	"todoList/models"
	"todoList/pkg/service"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	var taskInput models.Task
	if err := c.BindJSON(&taskInput); err != nil {
		logger.Error.Printf("[controllers.AddTask] error additing task: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid input data",
			"details": err.Error(),
		})
		return
	}

	newTask, err := service.AddTask(taskInput)
	if err != nil {
		logger.Error.Printf("[controllers.AddTask] error additing task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newTask)
}

func GetAllTasks(c *gin.Context) {
	tasks, err := service.GetAllTasks()
	if err != nil {
		logger.Error.Printf("[controllers.GetAllTasks] error getting all tasks: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetTaskByID] error getting task: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	task, err := service.GetTaskByID(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.GetTaskByID] error getting task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func EditTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.EditTaskByID] error editing task: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var taskInput models.Task
	if err := c.ShouldBindJSON(&taskInput); err != nil {
		logger.Error.Printf("[controllers.EditTaskByID] error editing task: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	updatedTask, err := service.EditTaskByID(uint(id), taskInput)
	if err != nil {
		logger.Error.Printf("[controllers.EditTaskByID] error editing task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func MarkTaskAsDoneByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.MarkTaskAsDoneByID] error marking task: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	task, err := service.MarkTaskAsDoneByID(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.MarkTaskAsDoneByID] error marking task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTaskByID] error deleating task: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = service.DeleteTaskByID(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTaskByID] error deleating task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}
