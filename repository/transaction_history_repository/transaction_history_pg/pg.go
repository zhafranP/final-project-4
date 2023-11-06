package transaction_history_pg

import (
	"context"
	"database/sql"
	"finalProject4/dto"
	"finalProject4/entity"
	"finalProject4/pkg/errs"
	"finalProject4/repository/transaction_history_repository"
	"strconv"
)

type transactionPG struct {
	db *sql.DB
}

const (
	findUserByIdQuery = `
	SELECT * FROM users WHERE id = $1
	`

	findProductByIdQuery = `
	SELECT * FROM products WHERE id = $1
	`

	createTransactionQuery = `
	INSERT INTO transaction_histories (product_id, user_id, quantity, total_price)
	VALUES ($1, $2, $3, $4)
	`

	updateUserBalanceQuery = `
	UPDATE users SET balance = balance - $1 WHERE id = $2
	`

	updateProductQuery = `
	UPDATE products SET stock = stock - $1 WHERE id = $2
	`

	updateCategoryQuery = `
	UPDATE categories set sold_product_amount = sold_product_amount + $1 WHERE id = $2
	`
	
)

func NewTransactionPG(db *sql.DB) transaction_history_repository.Repository {
	return &transactionPG{db: db}
}

func (transactionPG *transactionPG) CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.NewTransactionHistoryResponse, errs.Error) {
	var user entity.User
	var product entity.Product
	err := transactionPG.db.QueryRow(findUserByIdQuery, transactionPayload.UserID).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Password, &user.Role, &user.Balance, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}
	err = transactionPG.db.QueryRow(findProductByIdQuery, transactionPayload.ProductID).Scan(
		&product.ID, &product.Title, &product.Price, &product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("product with id: " + strconv.Itoa(transactionPayload.ProductID) + " not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	//Logic Check
	if (product.Stock - transactionPayload.Quantity < 5) {
		return nil, errs.NewBadRequest("quantity requested exceed product stock")
	}
	totalPrice := transactionPayload.Quantity * product.Price
	if (user.Balance < totalPrice) {
		return nil, errs.NewBadRequest("insufficient user balance, total price: " + strconv.Itoa(totalPrice))
	}

	// Run Query
	ctx := context.Background()
	tx, _ := transactionPG.db.BeginTx(ctx, nil)

	_, err = tx.ExecContext(ctx, updateUserBalanceQuery, totalPrice, user.ID)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	_, err = tx.ExecContext(ctx, updateProductQuery, transactionPayload.Quantity, product.ID)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	_, err = tx.ExecContext(ctx, updateCategoryQuery, transactionPayload.Quantity, product.CategoryID)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	_, err = tx.ExecContext(ctx, createTransactionQuery, product.ID, user.ID, transactionPayload.Quantity, totalPrice)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	res := dto.NewTransactionHistoryResponse{
		TotalPrice: totalPrice,
		Quantity: transactionPayload.Quantity,
		ProductTitle: product.Title,
	}
	return &res, nil
}
