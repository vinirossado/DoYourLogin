package repositories

import (
	"doYourLogin/source/domain/entities"
)

func CreateLog(log *entities.Logs, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Create(log).Error
}

//func FindLogByID(id uint, tx ...*TransactionalOperation) (*entities.Logs, error) {
//	log := &entities.Logs{}
//	return log, WithTransactions(tx).Where("id = ?", id).First(log).Error
//}
