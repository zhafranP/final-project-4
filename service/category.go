package service

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
	"finalProject4/pkg/helpers"
	"finalProject4/repository/category_repository"
)

type categoryService struct {
	categoryRepo category_repository.Repository
}

type CategoryService interface {
	CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error)
	GetCategories() (*dto.GetCategories, errs.Error)
}

func NewCategoryService(categoryRepo category_repository.Repository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (cs *categoryService) CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(categoryPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	res, err := cs.categoryRepo.CreateCategory(categoryPayload)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cs *categoryService) GetCategories() (*dto.GetCategories, errs.Error) {
	categories, err := cs.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}