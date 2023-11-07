package handler

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
	"finalProject4/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) transactionHandler {
	return transactionHandler{transactionService: transactionService}
}

func (th *transactionHandler) CreateTransaction(c *gin.Context) {
	var transactionPayload dto.NewTransactionHistoryRequest
	if err := c.ShouldBindJSON(&transactionPayload); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}

	jwtClaims := c.MustGet("user").(jwt.MapClaims)
	transactionPayload.UserID = int(jwtClaims["id"].(float64))

	res, err := th.transactionService.CreateTransaction(&transactionPayload)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (th *transactionHandler) GetTransactionUser(c *gin.Context) {
	jwtClaims := c.MustGet("user").(jwt.MapClaims)
	userId := int(jwtClaims["id"].(float64))
	res, err := th.transactionService.GetTransactionUser(userId)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (th *transactionHandler) GetTransactionAdmin(c *gin.Context) {
	res, err := th.transactionService.GetTransactionAdmin()
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, res)
}