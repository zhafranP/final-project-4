package dto

import "time"

type NewCategoryRequest struct {
	Type string `json:"type" validate:"required"`
}

type NewCategoryResponse struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_time"`
}

