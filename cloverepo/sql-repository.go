package cloverepo

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/mayron1806/go-clover-core/clovermodel"
)

type SQLRepository[T clovermodel.IModel] struct {
	db        *sql.DB
	tableName string
}

// NewSQLRepository creates a new instance of SQLRepository.
func NewSQLRepository[T clovermodel.IModel](db *sql.DB, tableName string) *SQLRepository[T] {
	return &SQLRepository[T]{
		db:        db,
		tableName: tableName,
	}
}

// validateEntity checks if the entity is valid before any operation.
func (r *SQLRepository[T]) validateEntity(entity T) error {
	if err := entity.Validate(); err != nil {
		return err
	}
	return nil
}

// Implementa o método de busca (Find) por chave primária
func (r *SQLRepository[T]) Find(key int64) (T, error) {
	var entity T
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", r.tableName)
	row := r.db.QueryRow(query, key)

	if err := scanRow(row, &entity); err != nil {
		return entity, err
	}
	return entity, nil
}

// FindAll retrieves all entities from the repository.
func (r *SQLRepository[T]) FindAll() ([]T, error) {
	query := fmt.Sprintf("SELECT * FROM %s", r.tableName)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entities []T
	if err = scanRows(rows, entities); err != nil {
		return nil, err
	}
	return entities, nil
}

// Implementa o método de criação (Create) com validação e hook BeforeCreate
func (r *SQLRepository[T]) Create(entity T) (T, error) {
	if err := r.validateEntity(entity); err != nil {
		return entity, err
	}
	if err := entity.BeforeCreate(); err != nil {
		return entity, err
	}

	fields, values, placeholders := extractFieldsAndValues(entity)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", r.tableName, strings.Join(fields, ","), strings.Join(placeholders, ","))

	_, err := r.db.Exec(query, values...)
	return entity, err
}

// Implementa o método de atualização (Update) com validação e hook BeforeUpdate
func (r *SQLRepository[T]) Update(entity T) (T, error) {
	// Validação e hook BeforeUpdate
	if err := r.validateEntity(entity); err != nil {
		return entity, err
	}
	if err := entity.BeforeUpdate(); err != nil {
		return entity, err
	}

	fields, values, _ := extractFieldsAndValues(entity)
	setClause := []string{}
	for _, field := range fields {
		setClause = append(setClause, fmt.Sprintf("%s = ?", field))
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", r.tableName, strings.Join(setClause, ","))
	_, err := r.db.Exec(query, append(values, reflect.ValueOf(entity).FieldByName("ID").Interface())...)
	return entity, err
}

// Implementa o método de exclusão (Delete) com validação e hook BeforeDelete
func (r *SQLRepository[T]) Delete(key int64, entity T) error {
	// Validação e hook BeforeDelete
	if err := entity.BeforeDelete(); err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.tableName)
	_, err := r.db.Exec(query, key)
	return err
}
