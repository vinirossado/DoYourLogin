package repositories

//func CreateLog(log *entities.Logs, tx ...*infra.TransactionalOperation) error {
//	return infra.WithTransaction(tx).Create(log).Error
//}
//
//func FindLogByID(id uint, tx ...*infra.TransactionalOperation) (*entities.Logs, error) {
//	log := &entities.Logs{}
//	return log, infra.WithTransaction(tx).Where("id = ?", id).First(log).Error
//}
