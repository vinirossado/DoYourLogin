package entities

func RetriveAll() []interface{} {
	return []interface{}{
		&Company{},
		&User{},
		&Logs{},
		&Book{},
		&Author{},
		&Category{},
	}
}
