package handler

import (
	"faq_sys_go/models"
	"faq_sys_go/repository"
	"faq_sys_go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FAQCategoryHandler struct {
	categoryRepo *repository.FAQCategoryRepository
}

func NewFAQCategoryHandler(categoryRepo *repository.FAQCategoryRepository) *FAQCategoryHandler {
	return &FAQCategoryHandler{categoryRepo: categoryRepo}
}

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *FAQCategoryHandler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category := &models.FAQCategory{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.categoryRepo.Create(category); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create category")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Category created successfully", category)
}

func (h *FAQCategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryRepo.FindAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}

func (h *FAQCategoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.categoryRepo.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Category not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Category retrieved successfully", category)
}

func (h *FAQCategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.categoryRepo.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Category not found")
		return
	}

	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}

	if err := h.categoryRepo.Update(category); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update category")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Category updated successfully", category)
}

func (h *FAQCategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	if err := h.categoryRepo.Delete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Category deleted successfully", nil)
}