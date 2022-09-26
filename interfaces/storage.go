package interfaces

import "github.com/fokurly/avito-tech-test-task/models"

type Storage interface {
	IncreaseBalance(client models.Client) (*float64, error)
	DecreaseBalance(client models.Client) (*float64, error)
	GetBalance(clientId int64) (*float64, error)
	GetAllClients() ([]models.Client, error)
}
