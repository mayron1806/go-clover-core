package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type IModel interface {
	Update(IModel) error
	Validate() error
	BeforeCreate() error
	BeforeUpdate() error
	BeforeDelete() error
}
type Model struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (m *Model) BeforeCreate() error {
	return nil
}

func (m *Model) BeforeUpdate() error {
	return nil
}

func (m *Model) BeforeDelete() error {
	return nil
}

func (m *Model) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
