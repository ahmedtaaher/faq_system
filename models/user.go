package models

import (
	"time"

	"gorm.io/gorm"
)

type UserType string

const (
	UserTypeCustomer UserType = "customer"
	UserTypeMerchant UserType = "merchant"
	UserTypeAdmin    UserType = "admin"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"` 
	UserType  UserType       `gorm:"type:varchar(20);not null" json:"user_type"`
	StoreID   *uint          `gorm:"index" json:"store_id,omitempty"`
	Store     *Store         `gorm:"foreignKey:StoreID" json:"store,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) IsMerchant() bool {
	return u.UserType == UserTypeMerchant
}

func (u *User) IsAdmin() bool {
	return u.UserType == UserTypeAdmin
}

func (u *User) IsCustomer() bool {
	return u.UserType == UserTypeCustomer
}