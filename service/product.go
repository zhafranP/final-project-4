package service

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
	"finalProject4/pkg/helpers"
	"finalProject4/repository/product_repository"
)

type ProductService interface {
	CreateProduct(p dto.NewProductRequest) (*dto.NewProductResponse, errs.Error)
	GetProducts() (*dto.GetProductResponse, errs.Error)
	UpdateProduct(p dto.UpdateProductRequest) (*dto.UpdateProductResponse, errs.Error)
	DeleteProduct(id int) (*dto.DeleteProductsResponse, errs.Error)
}

type productService struct {
	productRepo product_repository.Repository
}

func NewProductService(productRepo product_repository.Repository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (ps *productService) CreateProduct(p dto.NewProductRequest) (*dto.NewProductResponse, errs.Error) {

	validateErr := helpers.ValidateStruct(p)
	if validateErr != nil {
		return nil, validateErr
	}

	categoryExist, err := ps.productRepo.CategoryIdExist(p.CategoryID)
	if err != nil {
		return nil, err
	}

	if !categoryExist {
		return nil, errs.NewBadRequest("Category Not Exist")
	}

	resp, err := ps.productRepo.CreateProduct(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *productService) GetProducts() (*dto.GetProductResponse, errs.Error) {
	resp, err := ps.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *productService) UpdateProduct(p dto.UpdateProductRequest) (*dto.UpdateProductResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(p)
	if validateErr != nil {
		return nil, validateErr
	}

	categoryExist, err := ps.productRepo.CategoryIdExist(p.CategoryID)
	if err != nil {
		return nil, err
	}

	if !categoryExist {
		return nil, errs.NewBadRequest("Category Not Exist")
	}

	resp, err := ps.productRepo.UpdateProduct(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *productService) DeleteProduct(id int) (*dto.DeleteProductsResponse, errs.Error) {
	var resp dto.DeleteProductsResponse

	err := ps.productRepo.DeleteProducts(id)
	if err != nil {
		return nil, err
	}

	resp.Message = "Product Has Been Successfully Deleted"
	return &resp, nil
}
