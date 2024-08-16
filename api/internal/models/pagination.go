package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PaginationRequest struct {
	ContinuationToken string `query:"continuationToken" validate:"base64"`
	PageSize          int    `query:"pageSize" validate:"required,min=5,max=100"`
	OrderDesc         bool   `query:"orderDesc" validate:"bool"`
}

type PaginationResponse struct {
	ContinuationToken string `json:"continuationToken"`
	// TotalRecords      int
	// TotalPages        int
}

// Validate pagination parameters
func ValidatePagination(c *fiber.Ctx) error {
	var errors []*IError

	body := new(PaginationRequest)
	c.QueryParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
