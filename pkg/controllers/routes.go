package controllers

import (
	"fmt"
	"todoList/configs"

	"github.com/gin-gonic/gin"
)

func RunRoutes() error {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	taskG := r.Group("/tasks", checkUserAuthentication)
	{
		taskG.POST("/", AddTask)
		taskG.GET("/", GetAllTasks)
		taskG.GET("/:id", GetTaskByID)
		taskG.PUT("/:id/done", MarkTaskAsDoneByID)
		taskG.PUT("/:id", EditTaskByID)
		taskG.DELETE("/:id", DeleteTaskByID)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userG := r.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", EditUserByID)
		userG.DELETE("/:id", DeleteUserByID)
	}

	err := r.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))
	if err != nil {
		panic(err)
	}

	return nil
}
