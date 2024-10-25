package clover

import (
	"sync"

	"github.com/mayron1806/go-clover-core/db"
)

type Container struct {
	db *db.Database

	InitOnce sync.Once
}

func (c *Container) Database() *db.Database {
	return c.db
}
func NewContainer(db *db.Database) *Container {
	return &Container{
		db: db,
	}
}
