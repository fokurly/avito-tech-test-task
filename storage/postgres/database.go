package postgres

import (
	"fmt"
	"github.com/fokurly/avito-tech-test-task/models"
	"github.com/fokurly/avito-tech-test-task/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db     *sqlx.DB
	config models.DatabaseConfig
}

func NewDatabase() (*Db, error) {
	config := utils.ParseDatabaseConfigByKey("database_config", false)

	db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Dbname))
	if err != nil {
		return nil, err
	}

	// Убрать, потому что будут файлы sql, которые будут запускаться в докере перед основной программой
	db.MustExec(config.CreateTableString)

	return &Db{db: db, config: config}, nil
}
