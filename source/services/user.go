package services

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/exceptions"
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/logger"
	"doYourLogin/source/middlewares"
	"doYourLogin/source/repositories"
	"doYourLogin/source/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func FindUsers() *[]responses.UserResponse {
	users, err := repositories.FindUsers()

	if err != nil {
		exceptions.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get all users with error: %s", err),
			nil,
		)
	}

	usersResponse := []responses.UserResponse{}

	utils.Map(&users, &usersResponse)

	return &usersResponse
}

func FindUserById(id int) *responses.UserResponse {
	user, err := repositories.FindUserById(id)

	if err != nil {
		exceptions.ThrowNotFoundException(fmt.Sprintf("User with %d was not found", id),
			nil)
	}

	userResponse := responses.UserResponse{}

	utils.Map(user, &userResponse)

	return &userResponse
}

func CreateUser(request *requests.UserRequest) {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {

		exists := repositories.ExistsUserByUsernameOrEmail(request.Username, request.Email)

		if exists {
			return exceptions.BadRequestException(
				fmt.Sprintf("Username %s or Email %s already exists", request.Username, request.Email),
				nil,
			)
		}

		trace_logger.BuildLogger(fmt.Sprintf("Username %s already exists", request.Username), "/create-user", "CreateUser", http.StatusBadRequest, tx)
		hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		request.CompanyID = middlewares.TokenClaims.CompanyID
		request.Password = string(hash)

		var user = entities.User{}

		utils.Map(request, &user)

		if err := repositories.CreateUser(&user, tx); err != nil {
			return exceptions.InternalServerException(
				fmt.Sprintf("Error while trying to insert new User with error: %s", err),
				nil,
			)
		}

		companyDb, _ := repositories.FindCompanyById(request.CompanyID, tx)

		es := utils.InitEmailServer()

		err := es.SendEmail(user.Email, "Bando de fds", fmt.Sprintf("http://localhost:9025/activate-account/%s", companyDb.APIToken))

		if err != nil {
			return err
		}

		return nil
	})
}

func UpdateUser(request *requests.UserUpdateRequest, id int) {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		user, err := repositories.FindUserById(id)

		if err != nil {
			return exceptions.NotFoundException(
				fmt.Sprintf("User with id {%d} was not found", id),
				nil,
			)
		}

		utils.Map(request, user)

		if err := repositories.UpdateUser(user, tx); err != nil {
			return exceptions.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to update new user with error: %s", err),
				nil,
			)
		}
		return nil
	})
}

func DeleteUser(id int) {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		if err := repositories.DeleteUser(id, tx); err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return exceptions.NotFoundException(
					fmt.Sprintf("User with id {%d} not found", id),
					nil,
				)
			}

			return exceptions.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to delete new user with error: %s", err),
				nil,
			)
		}
		return nil
	})
}
