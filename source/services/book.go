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

func FindBooks() *[]entities.Book {
	books, err := repositories.FindBooks()

	if err != nil {
		exceptions.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get all books with error: %s", err),
			nil,
		)
	}

	return &books
}

func FindBookById(id int) *responses.BookResponse {
	book, err := repositories.FindBookById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exceptions.ThrowNotFoundException(
				fmt.Sprintf("Book with id {%d} was not found", id),
				nil,
			)
		}

		exceptions.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get book with id {%d} with error: %s", id, err),
			nil,
		)

	}

	bookResponse := responses.BookResponse{}

	utils.Map(book, &bookResponse)

	return &bookResponse
}

func CreateBook(request *requests.BookRequest) {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		var book = entities.Book{
			Name:        request.Name,
			ISBN:        request.ISBN,
			Description: request.Description,
			AuthorId:    request.AuthorId,
			Price:       request.Price,
			ReleaseYear: request.ReleaseYear,
		}

		if err := repositories.CreateBook(&book, tx); err != nil {
			return exceptions.InternalServerException(
				fmt.Sprintf("Error while trying to insert new Book with error: %s", err),
				nil,
			)
		}

		var categories []entities.Category
		if err := repositories.WithTransaction(tx).Where("id IN ?", request.Categories).Find(&categories).Error; err != nil {
			return err
		}

		return repositories.WithTransaction(tx).Model(&book).Association("Categories").Replace(&categories)
	})
}

func UpdateBook(request *requests.BookRequest, id int) {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		book, err := repositories.FindBookById(id)

		if err != nil {
			return exceptions.NotFoundException(
				fmt.Sprintf("Book with id {%d} was not found", id),
				nil,
			)
		}

		utils.Map(request, book)

		if err := repositories.UpdateBook(book, tx); err != nil {
			return exceptions.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to update book with error: %s", err),
				nil,
			)

		}
		return nil
	})
}

func DeleteBook(id int) {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		if err := repositories.DeleteBook(id, tx); err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return exceptions.NotFoundException(
					fmt.Sprintf("Book with id {%d} not found", id),
					nil,
				)
			}

			return exceptions.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to delete new book with error: %s", err),
				nil,
			)
		}
		return nil
	})
}
