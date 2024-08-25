package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

// Validate query parameters
func ValidateQueries(queryStruct interface{}) func(c *fiber.Ctx) error {

	var errors []*IError

	return func(c *fiber.Ctx) error {

		if err := c.QueryParser(queryStruct); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
		}

		if err := Validator.Struct(queryStruct); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err)
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
}

// Validate body parameters
func ValidateBody(bodyStuct interface{}) func(c *fiber.Ctx) error {

	var errors []*IError

	return func(c *fiber.Ctx) error {

		if err := c.BodyParser(bodyStuct); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
		}

		if err := Validator.Struct(bodyStuct); err != nil {
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
}
