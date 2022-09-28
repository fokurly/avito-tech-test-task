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
	//dbinfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, "5432", config.Dbname)
	query := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Dbname)
	logrus.Println(query)
	logrus.Println(config.Host)
	db, err := sqlx.Open("postgres", query)

	if err != nil {
		logrus.Println(err)
		return nil, err
	}

	return &Db{db: db, config: config}, nil
}
