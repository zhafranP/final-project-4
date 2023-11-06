package dto

type NewTransactionHistoryRequest struct {
	UserID    int
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,gt=0"`
}

type NewTransactionHistoryResponse struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionHistoryBill struct {
	Message         string                        `json:"message"`
	TransactionBill NewTransactionHistoryResponse `json:"transaction_bill"`
}
