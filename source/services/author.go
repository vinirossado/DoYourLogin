package services

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/exceptions"
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/repositories"
	"doYourLogin/source/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func FindAuthors() *[]responses.AuthorResponse {
	authors, err := repositories.FindAuthors()

	if err != nil {
		exceptions.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get all authors with error: %s", err),
			nil,
		)
	}

	authorsResponse := []responses.AuthorResponse{}

	utils.Map(&authors, &authorsResponse)

	return &authorsResponse
}

func FindAuthorById(id int) *responses.AuthorResponse {
	author, err := repositories.FindAuthorById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exceptions.ThrowNotFoundException(
				fmt.Sprintf("Author with id {%d} was not found", id),
				nil,
			)
		}

		exceptions.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get author with id {%d} with error: %s", id, err),
			nil,
		)

	}

	authorResponse := responses.AuthorResponse{}

	utils.Map(author, &authorResponse)

	return &authorResponse
}

func CreateAuthor(request *requests.AuthorRequest) {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		var author = entities.Author{}

		utils.Map(request, &author)

		if err := repositories.CreateAuthor(&author, tx); err != nil {
			return exceptions.InternalServerException(
				fmt.Sprintf("Error while trying to insert new Author with error: %s", err),
				nil,
			)
		}

		return nil
	})
}

func UpdateAuthor(request *requests.UserUpdateRequest, id int) {
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

func DeleteAuthor(id int) {
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
