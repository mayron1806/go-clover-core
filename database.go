package clover

import (
	"errors"

	"github.com/mayron1806/go-clover-core/cloverdb"
)

func (f *Clover) ConfigureDatabase(options cloverdb.DatabaseOptions) (*cloverdb.Database, error) {
	if f.db != nil {
		return f.db, nil
	}
	db, err := cloverdb.NewDatabase(&options)
	if err != nil {
		return nil, err
	}
	if err := db.Connect(); err != nil {
		return nil, err
	}
	f.db = db
	return db, nil
}
func (f *Clover) Database() (*cloverdb.Database, error) {
	if f.db != nil {
		return f.db, nil
	}
	return nil, errors.New("database not configured")
}
