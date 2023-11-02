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

	getCategoryQuery = `
	SELECT * FROM categories
	`
	getProductsByCategoryQuery = `
	SELECT * FROM products WHERE category_id = $1
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

func (categoryPG *categoryPG) GetCategories() (*dto.GetCategories, errs.Error) {
	var res dto.GetCategories
	var categories dto.GetCategoriesWithProducts
	var products dto.GetProductsCategory
	var temp int // for scan category_id

	categoryRows, err := categoryPG.db.Query(getCategoryQuery)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer categoryRows.Close()

	for categoryRows.Next() {
		err := categoryRows.Scan(
			&categories.ID, &categories.Type, &categories.SoldProductAmount, &categories.CreatedAt, &categories.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		productRows, err := categoryPG.db.Query(getProductsByCategoryQuery, categories.ID)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		for productRows.Next() {
			err = productRows.Scan(&products.ID, &products.Title, &products.Price, &products.Stock, &temp, &products.CreatedAt, &products.UpdatedAt)
			if err != nil {
				return nil, errs.NewInternalServerError(err.Error())
			}
			categories.Products = append(categories.Products, products)
		}
		productRows.Close()
		res.Data = append(res.Data, categories)
		categories.Products = nil
	}
	return &res, nil
}
