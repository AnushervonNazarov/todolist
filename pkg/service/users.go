package service

import (
	"errors"
	"fmt"
	"todoList/models"
	"todoList/pkg/repository"
	"todoList/utils"

	"gorm.io/gorm"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) error {
	_, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user.Password = utils.GenerateHash(user.Password)

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func EditUserByID(id uint, userInput models.User) (*models.User, error) {
	_, err := repository.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	userInput.ID = int(id)

	updatedUser, err := repository.EditUserByID(&userInput)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %v", err)
	}

	return updatedUser, nil
}

func DeleteUserByID(id uint) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	if err := repository.DeleteUserByID(&user); err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}

	return nil
}
