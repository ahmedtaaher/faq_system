package models
import (
	"time"

	"gorm.io/gorm"
)

type FAQCategory struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	FAQs        []FAQ          `gorm:"foreignKey:CategoryID" json:"faqs,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (FAQCategory) TableName() string {
	return "faq_categories"
}