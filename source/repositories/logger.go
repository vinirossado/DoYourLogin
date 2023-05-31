package repositories

import (
	"doYourLogin/source/domain/entities"
)

func CreateLog(log *entities.Logs, tx ...*TransactionalOperation) error {
	return WithTransaction(tx).Create(log).Error
}

//func FindLogByID(id uint, tx ...*TransactionalOperation) (*entities.Logs, error) {
//	log := &entities.Logs{}
//	return log, WithTransaction(tx).Where("id = ?", id).First(log).Error
//}
