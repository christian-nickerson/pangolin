package models

import (
	"time"

	"gorm.io/gorm"
)

type Database struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Name        string         `json:"name" gorm:"unique" validate:"required"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type DatabaseResponse struct {
	PaginationResponse
	Databases []Database
}
