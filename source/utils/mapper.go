package utils

import (
	"reflect"
)

func Map(source any, destination any) {
	sourceValue := reflect.ValueOf(source).Elem()
	destValue := reflect.ValueOf(destination).Elem()

	destType := destValue.Type()

	var maxIndex = 0

	if destValue.NumField() > sourceValue.NumField() {
		maxIndex = sourceValue.NumField()
	} else {
		maxIndex = destValue.NumField()
	}

	for i := 0; i < maxIndex; i++ {
		destField := destValue.Field(i)
		destFieldName := destType.Field(i).Name

		// Find corresponding source field by name
		sourceField := sourceValue.FieldByName(destFieldName)
		if !sourceField.IsValid() {
			continue
		}

		// Perform mapping from source to destination
		if destField.CanSet() && sourceField.IsValid() && destField.Type() == sourceField.Type() {
			destField.Set(sourceField)
		}
	}
}
