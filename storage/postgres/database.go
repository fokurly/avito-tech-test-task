package postgres

import (
	"fmt"
	"github.com/fokurly/avito-tech-test-task/models"
	"github.com/fokurly/avito-tech-test-task/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Db struct {
	db     *sqlx.DB
	config models.DatabaseConfig
}

func NewDatabase() (*Db, error) {
	config := utils.ParseDatabaseConfigByKey("database_config", false)
	query := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Dbname)
	logrus.Println(query)
	logrus.Println(config.CreateTableString)
	db, err := sqlx.Open("postgres", query)

	if err != nil {
		return nil, err
	}

	// Убрать, потому что будут файлы sql, которые будут запускаться в докере
	db.MustExec(config.CreateTableString)

	return &Db{db: db, config: config}, nil
}
