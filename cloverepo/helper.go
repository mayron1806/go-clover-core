package cloverepo

import (
	"database/sql"
	"reflect"
)

// extractFieldsAndValues extracts the fields and their values from an entity using reflection.
func extractFieldsAndValues[T any](entity T) (fields []string, values []interface{}, placeholders []string) {
	v := reflect.ValueOf(entity)
	t := reflect.TypeOf(entity)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// Ignore private fields
		if field.PkgPath != "" {
			continue
		}

		fields = append(fields, field.Name)
		values = append(values, value)
		placeholders = append(placeholders, "?")
	}

	return fields, values, placeholders
}

// scanRow scans a single row and fills the entity with its values.
func scanRow(row *sql.Row, entity interface{}) error {
	v := reflect.ValueOf(entity).Elem()
	fields := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fields[i] = v.Field(i).Addr().Interface()
	}
	return row.Scan(fields...)
}

// scanRows scans multiple rows and fills the slice of entities.
func scanRows(rows *sql.Rows, entities interface{}) error {
	v := reflect.ValueOf(entities).Elem()
	t := v.Type().Elem()

	for rows.Next() {
		entity := reflect.New(t).Elem()
		fields := make([]interface{}, entity.NumField())
		for i := 0; i < entity.NumField(); i++ {
			fields[i] = entity.Field(i).Addr().Interface()
		}
		rows.Scan(fields...)

		v.Set(reflect.Append(v, entity))
	}

	return nil
}
