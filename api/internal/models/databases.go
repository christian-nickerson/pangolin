package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Database struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	Name        string         `json:"name" gorm:"unique" validate:"required"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (d Database) GetID() uint { return d.ID }

type DatabaseResponse struct {
	Data []Database
	PaginationResponse
}

// Validate database body
func ValidateDatabase(c *fiber.Ctx) error {
	var errors []*IError
	var body Database

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	return c.Next()
}
