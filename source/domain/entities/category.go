package entities

type Category struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"column:name;unique"`
}
