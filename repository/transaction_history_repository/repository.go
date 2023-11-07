package transaction_history_repository

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
)

type Repository interface {
	CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.NewTransactionHistoryResponse, errs.Error)
	GetTransactionUser(userId int) (*[]dto.GetTransactionUser, errs.Error)
	GetTransactionAdmin() (*[]dto.GetTransactionAdmin, errs.Error)
}
