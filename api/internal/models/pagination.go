package models

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PaginationRequest struct {
	ContinuationToken string `query:"continuationToken" validate:"omitempty,base64"`
	PageSize          int    `query:"pageSize" validate:"required,min=5,max=100"`
	OrderDesc         bool   `query:"orderDesc" validate:"boolean"`
}

func (p *PaginationRequest) DecodeToken() (uint64, error) {
	decodedByte, err := base64.StdEncoding.DecodeString(p.ContinuationToken)
	decodedUint := binary.BigEndian.Uint64(decodedByte)
	return decodedUint, err
}

type PaginationResponse struct {
	ContinuationToken string `json:"continuationToken"`
	TotalRecords      int64  `json:"totalRecords"`
	TotalPages        int64  `json:"totalPages"`
}

// Validate pagination parameters
func ValidatePagination(c *fiber.Ctx) error {
	var errors []*IError
	var pagination PaginationRequest

	if err := c.QueryParser(&pagination); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	err := Validator.Struct(pagination)
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
