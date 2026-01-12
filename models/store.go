package models
import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	FAQs        []FAQ          `gorm:"foreignKey:StoreID" json:"faqs,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Store) TableName() string {
	return "stores"
}