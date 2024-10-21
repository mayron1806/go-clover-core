package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mayron1806/go-clover-core/logger"
)

type DatabaseOptions struct {
	Driver       string
	DSN          string
	MaxIdleTime  time.Duration
	MaxLifetime  time.Duration
	MaxIdleConns int
	MaxOpenConns int
}
type Database struct {
	dbInstance *sql.DB
	logger     *logger.Logger
	options    DatabaseOptions
}

func (d *Database) Connect() error {
	d.logger.Info("Connecting to database...")
	db, err := sql.Open(d.options.Driver, d.options.DSN)

	if d.options.MaxIdleConns > 0 {
		db.SetMaxIdleConns(d.options.MaxIdleConns)
	}
	if d.options.MaxOpenConns > 0 {
		db.SetMaxOpenConns(d.options.MaxOpenConns)
	}
	if d.options.MaxIdleTime > 0 {
		db.SetConnMaxIdleTime(d.options.MaxIdleTime)
	}
	if d.options.MaxLifetime > 0 {
		db.SetConnMaxLifetime(d.options.MaxLifetime)
	}

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
	d.logger.Info("Closing database...")
	return d.dbInstance.Close()
}
func NewDatabase(options *DatabaseOptions) (*Database, error) {
	if options == nil {
		return nil, errors.New("database options cannot be nil")
	}
	return &Database{
		options: *options,
		logger: logger.NewLogger(logger.LoggerOptions{
			Prefix: "DB",
		}),
	}, nil
}
