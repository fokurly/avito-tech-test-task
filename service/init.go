package service

import (
	"github.com/fokurly/avito-tech-test-task/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type clientBalanceService struct {
	storage interfaces.Storage
}

func NewClientBalanceService(storage interfaces.Storage) *clientBalanceService {
	return &clientBalanceService{storage: storage}
}

func (s *clientBalanceService) Run() error {
	eng := gin.Default()
	//eng.Use(s.authMiddleware)
	eng.GET("/get_client_balance", s.GetClientBalance)
	eng.GET("/get_all_clients", s.GetAllClients)
	eng.POST("/increase_client_balance", s.IncreaseClientBalance)
	eng.POST("/decrease_client_balance", s.DecreaseClientBalance)
	eng.POST("/transfer_money", s.TransferMoney)

	return eng.Run(":8070")

}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

	logrus.SetLevel(logrus.DebugLevel)
}
