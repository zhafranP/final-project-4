package handler

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
	"finalProject4/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) categoryHandler {
	return categoryHandler{categoryService: categoryService}
}

func (ch *categoryHandler) CreateCategory(c *gin.Context) {
	var category dto.NewCategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}

	res, err := ch.categoryService.CreateCategory(&category)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (ch *categoryHandler) GetCategories(c *gin.Context) {
	res, err := ch.categoryService.GetCategories()
	fmt.Println(res)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func (ch *categoryHandler) UpdateCategory(c *gin.Context) {
	var category dto.NewCategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}
	param := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(param)
	res, err := ch.categoryService.UpdateCategory(categoryId, &category)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ch *categoryHandler) DeleteCategory(c *gin.Context) {
	param := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(param)
	err := ch.categoryService.DeleteCategory(categoryId)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "category has been successfully deleted",
	})
}