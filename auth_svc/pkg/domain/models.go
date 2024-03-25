package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int32          `json:"id" gorm:"primarykey;auto_increment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	FirstName string         `json:"firstname" binding:"required"`
	LastName  string         `json:"lastname" binding:"required"`
	Email     string         `json:"email" binding:"required" gorm:"unique"`
	Phone     string         `json:"phone" binding:"required"`
}
