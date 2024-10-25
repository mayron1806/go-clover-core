package clover

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/mayron1806/go-clover-core/db"
	"github.com/mayron1806/go-clover-core/logger"
)

type IRepository interface {
	New(db *db.Database, logger *logger.Logger) *Repository
}
type Repository struct {
	db            *db.Database
	logger        *logger.Logger
	builderSelect *sqlbuilder.SelectBuilder
	builderInsert *sqlbuilder.InsertBuilder
	builderUpdate *sqlbuilder.UpdateBuilder
	builderDelete *sqlbuilder.DeleteBuilder
}

func NewRepository(db *db.Database, logger *logger.Logger) *Repository {
	driver := db.GetOptions().Driver
	var flavor sqlbuilder.Flavor
	switch driver {
	case "mysql":
		flavor = sqlbuilder.MySQL
	case "postgres":
		flavor = sqlbuilder.PostgreSQL
	case "sqlite":
		flavor = sqlbuilder.SQLite
	case "sqlserver":
		flavor = sqlbuilder.SQLServer
	case "cql":
		flavor = sqlbuilder.CQL
	case "clickhouse":
		flavor = sqlbuilder.ClickHouse
	case "presto":
		flavor = sqlbuilder.Presto
	case "oracle":
		flavor = sqlbuilder.Oracle
	case "informix":
		flavor = sqlbuilder.Informix
	default:
		panic(fmt.Errorf("unsupported db driver: %s", driver))
	}
	return &Repository{
		db:            db,
		logger:        logger,
		builderSelect: flavor.NewSelectBuilder(),
		builderInsert: flavor.NewInsertBuilder(),
		builderUpdate: flavor.NewUpdateBuilder(),
		builderDelete: flavor.NewDeleteBuilder(),
	}
}
