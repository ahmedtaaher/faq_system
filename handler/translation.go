package handler

import (
	"faq_sys_go/models"
	"faq_sys_go/repository"
	"faq_sys_go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TranslationHandler struct {
	translationRepo *repository.TranslationRepository
	faqRepo         *repository.FAQRepository
}

func NewTranslationHandler(translationRepo *repository.TranslationRepository, faqRepo *repository.FAQRepository) *TranslationHandler {
	return &TranslationHandler{
		translationRepo: translationRepo,
		faqRepo:         faqRepo,
	}
}

type CreateTranslationRequest struct {
	FAQID    uint   `json:"faq_id" binding:"required"`
	Language string `json:"language" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type UpdateTranslationRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (h *TranslationHandler) CreateTranslation(c *gin.Context) {
	var req CreateTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	faq, err := h.faqRepo.FindByID(req.FAQID)
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

	existing, _ := h.translationRepo.FindByFAQIDAndLanguage(req.FAQID, req.Language)
	if existing != nil {
		utils.ErrorResponse(c, http.StatusConflict, "Translation for this language already exists")
		return
	}

	translation := &models.FAQTranslation{
		FAQID:    req.FAQID,
		Language: req.Language,
		Question: req.Question,
		Answer:   req.Answer,
	}

	if err := h.translationRepo.Create(translation); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create translation")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Translation created successfully", translation)
}

func (h *TranslationHandler) GetTranslationsByFAQID(c *gin.Context) {
	faqID, err := strconv.ParseUint(c.Param("faq_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}

	translations, err := h.translationRepo.FindByFAQID(uint(faqID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch translations")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Translations retrieved successfully", translations)
}

func (h *TranslationHandler) UpdateTranslation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid translation ID")
		return
	}

	translation, err := h.translationRepo.FindByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Translation not found")
		return
	}

	faq, err := h.faqRepo.FindByID(translation.FAQID)
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

	var req UpdateTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if req.Question != "" {
		translation.Question = req.Question
	}
	if req.Answer != "" {
		translation.Answer = req.Answer
	}

	if err := h.translationRepo.Update(translation); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update translation")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Translation updated successfully", translation)
}

func (h *TranslationHandler) DeleteTranslation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid translation ID")
		return
	}

	if err := h.translationRepo.Delete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete translation")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Translation deleted successfully", nil)
}