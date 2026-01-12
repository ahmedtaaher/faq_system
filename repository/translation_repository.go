package repository

import (
	"faq_sys_go/models"

	"gorm.io/gorm"
)

type TranslationRepository struct {
	db *gorm.DB
}

func NewTranslationRepository(db *gorm.DB) *TranslationRepository {
	return &TranslationRepository{db: db}
}

func (r *TranslationRepository) Create(translation *models.FAQTranslation) error {
	return r.db.Create(translation).Error
}

func (r *TranslationRepository) FindByID(id uint) (*models.FAQTranslation, error) {
	var translation models.FAQTranslation
	err := r.db.Preload("FAQ").First(&translation, id).Error
	if err != nil {
		return nil, err
	}
	return &translation, nil
}

func (r *TranslationRepository) FindByFAQID(faqID uint) ([]models.FAQTranslation, error) {
	var translations []models.FAQTranslation
	err := r.db.Where("faq_id = ?", faqID).Find(&translations).Error
	return translations, err
}

func (r *TranslationRepository) FindByFAQIDAndLanguage(faqID uint, language string) (*models.FAQTranslation, error) {
	var translation models.FAQTranslation
	err := r.db.Where("faq_id = ? AND language = ?", faqID, language).First(&translation).Error
	if err != nil {
		return nil, err
	}
	return &translation, nil
}

func (r *TranslationRepository) Update(translation *models.FAQTranslation) error {
	return r.db.Save(translation).Error
}

func (r *TranslationRepository) Delete(id uint) error {
	return r.db.Delete(&models.FAQTranslation{}, id).Error
}

func (r *TranslationRepository) DeleteByFAQID(faqID uint) error {
	return r.db.Where("faq_id = ?", faqID).Delete(&models.FAQTranslation{}).Error
}