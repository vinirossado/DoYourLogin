package infra

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type TransactionalOperation struct {
	transaction *gorm.DB
}

func (to *TransactionalOperation) Commit() error {
	return to.transaction.Commit().Error
}

func (to *TransactionalOperation) Rollback() error {
	return to.transaction.Rollback().Error
}

func (to *TransactionalOperation) BeginTransaction() error {
	return to.transaction.Begin().Error
}

func UsingTransactional(fn func(*TransactionalOperation) error) {

	tx := db.Begin()

	to := &TransactionalOperation{transaction: tx}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			//utils.CreateLog(nil, nil)
			panic(r)
		} else if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else if err := fn(to); err != nil {
			tx.Rollback()
			panic(err)
		} else if err := to.Commit(); err != nil {
			tx.Rollback()
			panic(err)
		}
		tx.Commit()
	}()
}

func WithTransaction(tx []*TransactionalOperation) *gorm.DB {
	for _, t := range tx {
		if t != nil {
			return t.transaction
		}
	}
	return db
}
