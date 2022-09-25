package models

type DatabaseConfig struct {
	CreateTableString string `json:"create_database_string" validate:"required"`
	User              string `json:"user" validate:"required"`
	Password          string `json:"password" validate:"required"`
	Host              string `json:"host" validate:"required"`
	Dbname            string `json:"dbname" validate:"required"`
}
