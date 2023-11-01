package category_pg

import (
	"database/sql"
	"finalProject4/dto"
	"finalProject4/pkg/errs"
	"finalProject4/repository/category_repository"
)

const (
	createCategoryQuery = `
	INSERT INTO categories (type) VALUES ($1)
	RETURNING id, type, sold_product_amount, created_at
	`
)

type categoryPG struct {
	db *sql.DB
}

func NewCategoryPG(db *sql.DB) category_repository.Repository {
	return &categoryPG{db: db}
}

func (categoryPG *categoryPG) CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error) {
	var res dto.NewCategoryResponse
	err := categoryPG.db.QueryRow(createCategoryQuery, categoryPayload.Type).Scan(&res.ID, &res.Type, &res.SoldProductAmount, &res.CreatedAt)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &res, nil
}
