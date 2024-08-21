package models

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PaginationRequest struct {
	ContinuationToken string `query:"continuationToken" validate:"base64"`
	PageSize          int    `query:"pageSize" validate:"required,min=5,max=100"`
	OrderDesc         bool   `query:"orderDesc" validate:"bool"`
}

func (p *PaginationRequest) DecodeToken() (uint64, error) {
	decodedByte, err := base64.StdEncoding.DecodeString(p.ContinuationToken)
	decodedUint := binary.BigEndian.Uint64(decodedByte)
	return decodedUint, err
}

type PaginationResponse struct {
	ContinuationToken string `json:"continuationToken"`
	// TotalRecords      int
	// TotalPages        int
}

// Validate pagination parameters
func ValidatePagination(c *fiber.Ctx) error {
	var errors []*IError

	pagination := new(PaginationRequest)
	c.QueryParser(&pagination)

	err := Validator.Struct(pagination)
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
