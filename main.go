package main

import (
	"errors"
	"fmt"
	"todoList/configs"
	"todoList/db"
	"todoList/logger"
	"todoList/pkg/controllers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
	}

	err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}

	if err := logger.Init(); err != nil {
		panic(err)
	}

	if err := db.ConnectToDB(); err != nil {
		panic(err)
	}

	if err := db.Migrate(); err != nil {
		panic(err)
	}

	if err := controllers.RunRoutes(); err != nil {
		panic(err)
	}

	// cmd.Run()

}
