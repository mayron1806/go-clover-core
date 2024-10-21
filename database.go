package clover

import (
	"errors"

	"github.com/mayron1806/go-clover-core/db"
)

func (f *Clover) ConfigureDatabase(options db.DatabaseOptions) (*db.Database, error) {
	if f.db != nil {
		return f.db, nil
	}
	db, err := db.NewDatabase(&options)
	if err != nil {
		return nil, err
	}
	if err := db.Connect(); err != nil {
		return nil, err
	}
	f.db = db
	return db, nil
}
func (f *Clover) Database() (*db.Database, error) {
	if f.db != nil {
		return f.db, nil
	}
	return nil, errors.New("database not configured")
}
