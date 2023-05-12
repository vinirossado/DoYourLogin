package repository

import (
	"abrigos/source/configuration"
	"abrigos/source/domain/exception"
	"fmt"
	"log"
	"math"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlserver.Open(configuration.DATABASE_SOURCE.ValueAsString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("An error ocurred during try to connect a database ", err)
	}

	AutoMigrate()
}

func GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@%s",
		configuration.DATABASE_USERNAME.ValueAsString(),
		configuration.DATABASE_PASSWORD.ValueAsString(),
		configuration.DATABASE_SOURCE.ValueAsString(),
	)
}

// Transaction

type TransactionalOperation struct {
	transaction *gorm.DB
}

func UsingTransactional(fn func(*TransactionalOperation) error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		return fn(&TransactionalOperation{transaction: tx})
	})
	if err != nil {
		except, ok := err.(*exception.HttpException)
		if ok {
			panic(except)
		}

		panic(exception.InternalServerException(err.Error()))
	}
}

func WithTransaction(tx []*TransactionalOperation) *gorm.DB {
	for idx, t := range tx {
		if t != nil {
			return tx[idx].transaction
		}
	}

	return db
}

//Pagination

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	Sort       string      `json:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id asc"
	}
	return p.Sort
}

func PaginateScope(model interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(model).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
