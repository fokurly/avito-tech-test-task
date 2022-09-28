package service

import (
	"fmt"
	"github.com/fokurly/avito-tech-test-task/models"
	"github.com/fokurly/avito-tech-test-task/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

const (
	base10  = 10
	bitSize = 64
)

func (s *clientBalanceService) GetClientBalance(ctx *gin.Context) {
	strId := ctx.Query("client_id")
	if strId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("empty client_id"))
		return
	}

	clientId, err := strconv.ParseInt(strId, base10, bitSize)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("failed to parse client_id. error: %v", err))
		return
	}

	balance, err := s.storage.GetBalance(clientId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("failed to get client balance. error: %v", err))
		return
	}

	currency := ctx.Query("currency")
	if currency != "" {
		newBalance, err := utils.ConvertBalanceToAnotherCurrency(*balance, currency)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not convert to another currency. error: %v", err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{fmt.Sprintf("client_balance (%s)", currency): newBalance})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"client_balance (RUB)": balance})
}

func (s *clientBalanceService) GetAllClients(ctx *gin.Context) {
	employees, err := s.storage.GetAllClients()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not get all employees from storage. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, employees)
}

func (s *clientBalanceService) IncreaseClientBalance(ctx *gin.Context) {
	var client models.Client
	if err := ctx.MustBindWith(&client, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. error: %v", err))
		return
	}

	currentBalance, err := s.storage.IncreaseBalance(client)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not increase client's balance. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("money were transferred to the account with id %d. current balance is %f", client.Id, *currentBalance))
}

func (s *clientBalanceService) DecreaseClientBalance(ctx *gin.Context) {
	var client models.Client
	if err := ctx.MustBindWith(&client, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. error: %v", err))
		return
	}

	currentBalance, err := s.storage.DecreaseBalance(client)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not decrease client's balance. error: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("money was withdrawn from the account with id %d. current balance is %f", client.Id, *currentBalance))
}

func (s *clientBalanceService) TransferMoney(ctx *gin.Context) {
	var transfer models.Transfer
	if err := ctx.MustBindWith(&transfer, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. error: %v", err))
		return
	}

	senderClient := &models.Client{Id: transfer.SenderId, Money: transfer.Amount}
	receiverClient := &models.Client{Id: transfer.SenderId, Money: transfer.Amount}

	currentSenderBalance, err := s.storage.DecreaseBalance(*senderClient)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not decrease client's balance. error: %v. transfer declined.", err))
		return
	}

	currentReceiverBalance, err := s.storage.IncreaseBalance(*receiverClient)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("could not increase client's balance. error: %v. transfer declined.", err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("transfer done. clients %d balance %f; clients %d balance, %f",
		senderClient.Id, *currentSenderBalance, receiverClient.Id, *currentReceiverBalance))
}
