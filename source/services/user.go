package service

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/exception"
	"doYourLogin/source/domain/request"
	"doYourLogin/source/domain/response"
	"doYourLogin/source/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FindUsers() []response.UserResponse {
	users, err := repository.FindUsers()

	if err != nil {
		exception.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get all users with error: %s", err),
		)
	}

	usersResponse := []response.UserResponse{}

	for _, user := range users {
		usersResponse = append(usersResponse, *MapToUserResponse(&user))
	}

	return usersResponse
}

func FindUserById(id int) *response.UserResponse {
	user, err := repository.FindUserById(id)

	if err != nil {
		exception.ThrowNotFoundException(fmt.Sprintf("User with %d was not found", id))
	}

	return MapToUserResponse(user)
}

func CreateUser(request *request.UserRequest) {

	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		exists := repository.ExistsUserByUsername(request.Username)

		if exists {
			return exception.BadRequestException(
				fmt.Sprintf("Username %s already exists", request.Username),
			)
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

		user := entities.User{
			Name:     request.Name,
			Username: request.Username,
			Password: string(hash),
			Role:     request.Role,
			Email:    request.Email,
			Address:  request.Address,
			Phone:    request.Phone,
			About:    request.About,
			Image:    request.Image,
		}

		if err := repository.CreateUser(&user, tx); err != nil {
			return exception.InternalServerException(
				fmt.Sprintf("Error while trying to insert new User with error: %s", err),
			)
		}

		return nil
	})
}

func UpdateUser(request *request.UserRequest, id int) {
	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		user, err := repository.FindUserById(id)

		if err != nil {
			return exception.NotFoundException(
				fmt.Sprintf("User with id {%d} was not found", id),
			)
		}

		user.Name = request.Name
		user.Username = request.Username
		user.Password = request.Password
		user.Role = request.Role
		user.Email = request.Email
		user.Address = request.Address
		user.Phone = request.Phone
		user.About = request.About
		user.Image = request.Image

		if err := repository.UpdateUser(user, tx); err != nil {
			return exception.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to update new user with error: %s", err))
		}
		return nil
	})
}

func DeleteUser(id int) {
	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		if err := repository.DeleteUser(id, tx); err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return exception.NotFoundException(
					fmt.Sprintf("User with id {%d} not found", id))
			}

			return exception.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to delete new user with error: %s", err))
		}
		return nil
	})
}

func MapToUserResponse(user *entities.User) (response *response.UserResponse) {

	return &response.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		About:    user.About,
		Image:    user.Image,
	}
}
