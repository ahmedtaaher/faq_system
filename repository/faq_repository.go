package repository

import (
	"faq_sys_go/models"

	"gorm.io/gorm"
)

type FAQRepository struct {
	db *gorm.DB
}

func NewFAQRepository(db *gorm.DB) *FAQRepository {
	return &FAQRepository{db: db}
}

func (r *FAQRepository) Create(faq *models.FAQ) error {
	return r.db.Create(faq).Error
}

func (r *FAQRepository) FindAll() ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.db.Preload("Category").Preload("Store").Preload("Translations").Find(&faqs).Error
	return faqs, err
}

func (r *FAQRepository) FindByID(id uint) (*models.FAQ, error) {
	var faq models.FAQ
	err := r.db.Preload("Category").Preload("Store").Preload("Translations").First(&faq, id).Error
	if err != nil {
		return nil, err
	}
	return &faq, nil
}

func (r *FAQRepository) FindByStoreID(storeID uint) ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.db.Preload("Category").Preload("Translations").
		Where("store_id = ? OR is_global = ?", storeID, true).
		Find(&faqs).Error
	return faqs, err
}

func (r *FAQRepository) FindGlobalFAQs() ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.db.Preload("Category").Preload("Translations").
		Where("is_global = ?", true).
		Find(&faqs).Error
	return faqs, err
}

func (r *FAQRepository) FindByCategoryAndStore(categoryID uint, storeID *uint) ([]models.FAQ, error) {
	var faqs []models.FAQ
	query := r.db.Preload("Translations").Where("category_id = ?", categoryID)
	
	if storeID != nil {
		query = query.Where("(store_id = ? OR is_global = ?)", *storeID, true)
	} else {
		query = query.Where("is_global = ?", true)
	}
	
	err := query.Find(&faqs).Error
	return faqs, err
}

func (r *FAQRepository) Update(faq *models.FAQ) error {
	return r.db.Save(faq).Error
}

func (r *FAQRepository) Delete(id uint) error {
	return r.db.Delete(&models.FAQ{}, id).Error
}