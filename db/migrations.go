package db

import (
	"todoList/models"
)

func Migrate() error {
	err := dbConn.AutoMigrate(models.Task{},
		models.User{})
	if err != nil {
		return err
	}
	return nil
}
