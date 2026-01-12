package models
import (
	"time"

	"gorm.io/gorm"
)

type FAQTranslation struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FAQID     uint           `gorm:"not null;uniqueIndex:idx_faq_language" json:"faq_id"`
	FAQ       FAQ            `gorm:"foreignKey:FAQID" json:"faq,omitempty"`
	Language  string         `gorm:"type:varchar(10);not null;uniqueIndex:idx_faq_language" json:"language"`
	Question  string         `gorm:"not null" json:"question"`
	Answer    string         `gorm:"type:text;not null" json:"answer"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (FAQTranslation) TableName() string {
	return "faq_translations"
}
