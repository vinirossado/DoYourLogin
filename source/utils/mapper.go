package utils

import (
	"fmt"
	"reflect"
)

func Map(source any, destination any) error {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(destination)

	sourceType := sourceValue.Type()
	destType := destValue.Type()

	if sourceType.Kind() == reflect.Slice && destType.Kind() == reflect.Slice {
		if sourceType.Elem().Kind() != reflect.Struct || destType.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("slice elements must be structs")
		}

		for i := 0; i < sourceValue.Len(); i++ {
			sourceElem := sourceValue.Index(i)
			destElem := reflect.New(destType.Elem()).Elem()

			err := Map(sourceElem.Interface(), destElem.Interface())

			if err != nil {
				return err
			}
			destValue = reflect.Append(destValue, destElem)
		}

		reflect.ValueOf(destination).Elem().Set(destValue)
	} else if sourceType.Kind() == reflect.Struct && destType.Kind() == reflect.Struct {
		if sourceType != destType {
			return fmt.Errorf("source and destination must have the same type")
		}

		for i := 0; i < sourceType.NumField(); i++ {
			sourceField := sourceType.Field(i)
			sourceFieldValue := sourceValue.Field(i)

			destField, found := destType.FieldByName(sourceField.Name)
			if found {
				destFieldValue := destValue.FieldByName(destField.Name)
				if destFieldValue.CanSet() && sourceFieldValue.Type() == destFieldValue.Type() {
					destFieldValue.Set(sourceFieldValue)
				}
			}

		}

	} else {
		return fmt.Errorf("Source and Destination must either be both structs or both slices")
	}
	return nil
}
