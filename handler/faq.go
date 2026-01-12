package handler

import (
	"faq_sys_go/models"
	"faq_sys_go/repository"
	"faq_sys_go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FAQHandler struct {
	faqRepo *repository.FAQRepository
}

func NewFAQHandler(faqRepo *repository.FAQRepository) *FAQHandler {
	return &FAQHandler{faqRepo: faqRepo}
}

type CreateFAQRequest struct {
	CategoryID uint   `json:"category_id" binding:"required"`
	Question   string `json:"question" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
	IsGlobal   bool   `json:"is_global"`
}

type UpdateFAQRequest struct {
	CategoryID uint   `json:"category_id"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	IsGlobal   bool   `json:"is_global"`
}

func (h *FAQHandler) CreateFAQ(c *gin.Context) {
	var req CreateFAQRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userType := c.GetString("userType")
	
	faq := &models.FAQ{
		CategoryID: req.CategoryID,
		Question:   req.Question,
		Answer:     req.Answer,
	}

	if userType == "admin" {
		faq.IsGlobal = req.IsGlobal
	} else if userType == "merchant" {
		storeID, exists := c.Get("storeID")
		if !exists || storeID == nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Merchant must have an associated store")
			return
		}
		storeIDUint := storeID.(*uint)
		faq.StoreID = storeIDUint
		faq.IsGlobal = false
	}

	if err := h.faqRepo.Create(faq); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create FAQ")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "FAQ created successfully", faq)
}

func (h *FAQHandler) GetAllFAQs(c *gin.Context) {
	userType := c.GetString("userType")
	language := c.Query("language")

	var faqs []models.FAQ
	var err error

	if userType == "admin" {
		faqs, err = h.faqRepo.FindAll()
	} else if userType == "merchant" {
		storeID, exists := c.Get("storeID")
		if !exists || storeID == nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Merchant must have an associated store")
			return
		}
		faqs, err = h.faqRepo.FindByStoreID(*storeID.(*uint))
	} else {
		faqs, err = h.faqRepo.FindGlobalFAQs()
	}

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch FAQs")
		return
	}

	if language != "" {
		for i := range faqs {
			filtered := []models.FAQTranslation{}
			for _, t := range faqs[i].Translations {
				if t.Language == language {
					filtered = append(filtered, t)
				}
			}
			faqs[i].Translations = filtered
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "FAQs retrieved successfully", faqs)
}

func (h *FAQHandler) GetFAQByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}

	faq, err := h.faqRepo.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "FAQ not found")
		return
	}

	userType := c.GetString("userType")
	if userType == "customer" && !faq.IsGlobal {
		utils.ErrorResponse(c, http.StatusForbidden, "Access denied")
		return
	}

	if userType == "merchant" && !faq.IsGlobal {
		storeID := c.MustGet("storeID").(*uint)
		if faq.StoreID == nil || *faq.StoreID != *storeID {
			utils.ErrorResponse(c, http.StatusForbidden, "Access denied")
			return
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "FAQ retrieved successfully", faq)
}

func (h *FAQHandler) UpdateFAQ(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}

	faq, err := h.faqRepo.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "FAQ not found")
		return
	}

	userType := c.GetString("userType")
	if userType == "merchant" {
		storeID := c.MustGet("storeID").(*uint)
		if faq.StoreID == nil || *faq.StoreID != *storeID {
			utils.ErrorResponse(c, http.StatusForbidden, "Access denied")
			return
		}
	}

	var req UpdateFAQRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if req.CategoryID != 0 {
		faq.CategoryID = req.CategoryID
	}
	if req.Question != "" {
		faq.Question = req.Question
	}
	if req.Answer != "" {
		faq.Answer = req.Answer
	}
	if userType == "admin" {
		faq.IsGlobal = req.IsGlobal
	}

	if err := h.faqRepo.Update(faq); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update FAQ")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "FAQ updated successfully", faq)
}

func (h *FAQHandler) DeleteFAQ(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}

	faq, err := h.faqRepo.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "FAQ not found")
		return
	}

	userType := c.GetString("userType")
	if userType == "merchant" {
		storeID := c.MustGet("storeID").(*uint)
		if faq.StoreID == nil || *faq.StoreID != *storeID {
			utils.ErrorResponse(c, http.StatusForbidden, "Access denied")
			return
		}
	}

	if err := h.faqRepo.Delete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete FAQ")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "FAQ deleted successfully", nil)
}