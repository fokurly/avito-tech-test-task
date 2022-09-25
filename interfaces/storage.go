package interfaces

type Storage interface {
	IncreaseBalance(money float64) error
	DecreaseBalance(money float64) error
	GetBalance(clientId int64) (*float64, error)
}
