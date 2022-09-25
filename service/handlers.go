package service

import (
	"fmt"
	"github.com/fokurly/avito-tech-test-task/models"
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

	// Может быть возвращать целиком клиента?
	ctx.JSON(http.StatusOK, gin.H{"client_balance": balance})
}

// Вопрос. Как добавлять клиентов? По условию сказано, что данные о балансе появляются при первом пополнении баланса, поэтому примем,
// что они уже как будто существуют, но их баланс равен 0.
// Или сделать добавление клиента??
func (s *clientBalanceService) IncreaseClientBalance(ctx *gin.Context) {
	var transfer models.Transfer
	if err := ctx.MustBindWith(&transfer, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not validate body. error: %v", err))
		return
	}

	if err := s.storage.IncreaseBalance()
}

func (s *clientBalanceService) DecreaseClientBalance(ctx *gin.Context) {

}

// Принимает json структуру с 3мя полями. Кому, от кого, сумма
func (s *clientBalanceService) TransferMoney(ctx *gin.Context) {

}
