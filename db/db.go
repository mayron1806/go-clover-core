package db

import (
	"database/sql"

	"github.com/mayron1806/go-clover-core/logging"
)

type Database struct {
	dbInstance *sql.DB
	logger     *logging.Logger
}

func (d *Database) Connect(driver, dsn string) error {
	d.logger.Info("Connecting to database...")
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		d.logger.Info("Failed to connect to database")
		return err
	}
	d.logger.Info("Connected to database")
	d.dbInstance = db
	return nil
}
func (d *Database) GetDB() *sql.DB {
	return d.dbInstance
}
func (d *Database) Close() error {
	return d.dbInstance.Close()
}
func NewDatabase() *Database {
	return &Database{
		logger: logging.NewLogger("Database"),
	}
}
