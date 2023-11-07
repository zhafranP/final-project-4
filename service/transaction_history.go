package service

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
	"finalProject4/pkg/helpers"
	"finalProject4/repository/transaction_history_repository"
)

type transactionService struct {
	transactionRepo transaction_history_repository.Repository
}

type TransactionService interface {
	CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.TransactionHistoryBill, errs.Error)
	GetTransactionUser(userId int) (*[]dto.GetTransactionUser, errs.Error)
	GetTransactionAdmin() (*[]dto.GetTransactionAdmin, errs.Error)
}

func NewTransactionService(transactionRepo transaction_history_repository.Repository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (ts *transactionService) CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.TransactionHistoryBill, errs.Error) {
	validateErr := helpers.ValidateStruct(transactionPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	transaction, err := ts.transactionRepo.CreateTransaction(transactionPayload)
	if err != nil {
		return nil, err
	}
	res := dto.TransactionHistoryBill{
		Message:         "You have successfully purchased the product",
		TransactionBill: *transaction,
	}
	return &res, nil
}

func (ts *transactionService) GetTransactionUser(userId int) (*[]dto.GetTransactionUser, errs.Error) {
	transactions, err := ts.transactionRepo.GetTransactionUser(userId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ts *transactionService) GetTransactionAdmin() (*[]dto.GetTransactionAdmin, errs.Error) {
	transactions, err := ts.transactionRepo.GetTransactionAdmin()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
