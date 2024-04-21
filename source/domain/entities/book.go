package entities

import "time"

type Book struct {
	ID          uint    `gorm:"primarykey"`
	AuthorId    uint    `gorm:"column:author_id"`
	Name        string  `gorm:"column:name"`
	ISBN        string  `gorm:"column:isbn;unique"`
	Active      bool    `gorm:"column:active"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
	ReleaseYear int     `gorm:"column:release_year"`

	Author     Author     `gorm:"foreignKey:AuthorId;references:ID"`
	Categories []Category `gorm:"many2many:book_categories;"`
	CreatedAt  time.Time  `gorm:"created_at"`
	UpdatedAt  time.Time  `gorm:"updated_at"`
	DeletedAt  time.Time  `gorm:"index"`
}
