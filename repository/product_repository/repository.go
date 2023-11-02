package product_repository

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
)

type Repository interface {
	CreateProduct(p dto.NewProductRequest) (*dto.NewProductResponse, errs.Error)
	UpdateProduct(p dto.UpdateProductRequest) (*dto.UpdateProductResponse, errs.Error)
	GetProducts() (*dto.GetProductResponse, errs.Error)
	DeleteProducts(id int) errs.Error
	CategoryIdExist(id int) (bool, errs.Error)
}
