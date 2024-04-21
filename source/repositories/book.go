package repositories

import (
	"doYourLogin/source/domain/entities"
)

func FindBooks(tx ...*TransactionalOperation) ([]entities.Book, error) {
	var books []entities.Book
	return books, WithTransactions(tx).Preload("Categories").Find(&books).Error
}

func FindBookById(id int, tx ...*TransactionalOperation) (*entities.Book, error) {
	book := entities.Book{}
	return &book, WithTransactions(tx).Where("id = ?", id).First(&book).Error
}

func CreateBook(book *entities.Book, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Create(book).Error
}

func UpdateBook(book *entities.Book, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Updates(&book).Error
}

func DeleteBook(id int, tx ...*TransactionalOperation) error {
	//dbResult := WithTransactions(tx).Delete(&entities.Book{Model: gorm.Model{ID: uint(id)}})
	//
	//if dbResult.RowsAffected < 1 {
	//	return gorm.ErrRecordNotFound
	//}
	//
	//return dbResult.Error
	return nil
}
