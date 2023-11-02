package category_repository

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
)

type Repository interface {
	CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error)
	GetCategories() (*dto.GetCategories, errs.Error)
}