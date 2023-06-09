package repositories

import (
	"doYourLogin/source/configuration"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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

//func GetConnectionString() string {
//
//	viper.SetConfigFile(".env")
//	viper.SetDefault("DATABASE_SOURCE", "localhost")
//
//	err := viper.ReadInConfig()
//	if err != nil {
//		log.Fatalf("Error reading config file: %s", err)
//	}
//
//	databaseSource := viper.GetString("DATABASE_SOURCE")
//	databaseUsername := viper.GetString("DATABASE_USERNAME")
//	databasePassword := viper.GetString("DATABASE_PASSWORD")
//
//	return fmt.Sprintf(
//		"%s:%s@%s",
//		databaseUsername,
//		databasePassword,
//		databaseSource,
//	)
//}
//
//// Transaction

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

//Pagination
//
//type Pagination struct {
//	Limit      int         `json:"limit"`
//	Page       int         `json:"page"`
//	Sort       string      `json:"sort"`
//	TotalRows  int64       `json:"total_rows"`
//	TotalPages int         `json:"total_pages"`
//	Rows       interface{} `json:"rows"`
//}
//
//func (p *Pagination) GetOffset() int {
//	return (p.GetPage() - 1) * p.GetLimit()
//}
//
//func (p *Pagination) GetLimit() int {
//	if p.Limit <= 0 {
//		p.Limit = 10
//	}
//
//	return p.Limit
//}
//
//func (p *Pagination) GetPage() int {
//	if p.Page <= 0 {
//		p.Page = 1
//	}
//	return p.Page
//}
//
//func (p *Pagination) GetSort() string {
//	if p.Sort == "" {
//		p.Sort = "Id asc"
//	}
//	return p.Sort
//}
//
//func PaginateScope(model interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
//	var totalRows int64
//	db.Model(model).Count(&totalRows)
//
//	pagination.TotalRows = totalRows
//	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
//
//	return func(db *gorm.DB) *gorm.DB {
//		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
//	}
//}
