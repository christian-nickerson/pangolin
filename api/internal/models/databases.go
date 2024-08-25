package models

import (
	"time"

	"gorm.io/gorm"
)

type Database struct {
	ID          uint64         `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Name        string         `json:"name" gorm:"unique" validate:"required"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (d Database) GetID() uint64 { return d.ID }

type DatabaseResponse struct {
	Data []Database
	PaginationResponse
}
