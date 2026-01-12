package models
import (
	"time"

	"gorm.io/gorm"
)

type FAQ struct {
	ID           uint             `gorm:"primaryKey" json:"id"`
	CategoryID   uint             `gorm:"not null;index" json:"category_id"`
	Category     FAQCategory      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	StoreID      *uint            `gorm:"index" json:"store_id,omitempty"` 
	Store        *Store           `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	IsGlobal     bool             `gorm:"default:false;index" json:"is_global"`
	Question     string           `gorm:"not null" json:"question"` 
	Answer       string           `gorm:"type:text;not null" json:"answer"` 
	Translations []FAQTranslation `gorm:"foreignKey:FAQID" json:"translations,omitempty"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	DeletedAt    gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (FAQ) TableName() string {
	return "faqs"
}