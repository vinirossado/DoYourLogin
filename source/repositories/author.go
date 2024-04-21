package repositories

import (
	"doYourLogin/source/domain/entities"
)

func FindAuthors(tx ...*TransactionalOperation) ([]entities.Author, error) {
	var authors []entities.Author
	return authors, WithTransactions(tx).Find(&authors).Error
}

func FindAuthorById(id int, tx ...*TransactionalOperation) (*entities.Author, error) {
	author := entities.Author{}
	return &author, WithTransactions(tx).Where("id = ?", id).First(&author).Error
}

func CreateAuthor(author *entities.Author, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Create(author).Error
}

func UpdateAuthor(author *entities.Author, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Updates(&author).Error
}

func DeleteAuthor(id int, tx ...*TransactionalOperation) error {
	//dbResult := WithTransactions(tx).Delete(&entities.Author{Model: gorm.Model{ID: uint(id)}})
	//
	//if dbResult.RowsAffected < 1 {
	//	return gorm.ErrRecordNotFound
	//}
	//
	//return dbResult.Error
	return nil
}
