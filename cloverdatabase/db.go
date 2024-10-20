package cloverdatabase

import (
	"database/sql"

	"github.com/mayron1806/go-clover-core/cloverlog"
)

type Database struct {
	dbInstance *sql.DB
	logger     *cloverlog.Logger
}

func (d *Database) Connect(driver, dsn string) error {
	d.logger.Info("Connecting to database...")
	db, err := sql.Open(driver, dsn)
	if err != nil {
		d.logger.Errorf("Failed to connect to database: %v", err.Error())
		return err
	}
	if err := db.Ping(); err != nil {
		d.logger.Errorf("Failed to ping database: %v", err.Error())
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
		logger: cloverlog.NewLogger(cloverlog.LoggerOptions{
			Prefix: "DB",
		}),
	}
}
