package repository

import (
	"faq_sys_go/models"

	"gorm.io/gorm"
)

type FAQCategoryRepository struct {
	db *gorm.DB
}

func NewFAQCategoryRepository(db *gorm.DB) *FAQCategoryRepository {
	return &FAQCategoryRepository{db: db}
}

func (r *FAQCategoryRepository) Create(category *models.FAQCategory) error {
	return r.db.Create(category).Error
}

func (r *FAQCategoryRepository) FindAll() ([]models.FAQCategory, error) {
	var categories []models.FAQCategory
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *FAQCategoryRepository) FindByID(id uint) (*models.FAQCategory, error) {
	var category models.FAQCategory
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *FAQCategoryRepository) Update(category *models.FAQCategory) error {
	return r.db.Save(category).Error
}

func (r *FAQCategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.FAQCategory{}, id).Error
}