package v1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/christian-nickerson/pangolin/api/internal/engines/databases"
	"github.com/christian-nickerson/pangolin/api/internal/models"
)

// add database routes to a fiber app
func AddDatabaseRoutes(route fiber.Router) {
	group := route.Group("/databases")

	group.Get("", models.ValidateQueries(&models.PaginationRequest{}), GetDatabases())
	group.Get("/:id", GetDatabase())
	group.Post("", models.ValidateBody(&models.Database{}), CreateDatabase())
	group.Patch("/:id", models.ValidateBody(&models.Database{}), UpdateDatabase())
	group.Delete("/:id", DeleteDatabase())
}

// return paginated database records
func GetDatabases() func(c *fiber.Ctx) error {

	var prq models.PaginationRequest
	var prp models.PaginationResponse
	var records []models.Database

	return func(c *fiber.Ctx) error {

		if err := c.QueryParser(&prq); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
		}

		if result := databases.DB.Scopes(
			databases.Paginate(records, &prq, &prp, databases.DB),
		).Find(&records); result.RowsAffected == 0 {
			return c.Status(fiber.StatusNoContent).SendString("No records found")
		}

		records, prp.ContinuationToken = databases.PaginatedResponse(records, prq.PageSize)
		return c.Status(fiber.StatusOK).JSON(
			models.DatabaseResponse{
				Data:               records,
				PaginationResponse: prp,
			},
		)
	}
}

// return a specific database record
func GetDatabase() func(c *fiber.Ctx) error {

	var record models.Database

	return func(c *fiber.Ctx) error {

		id := c.Params("id")
		result := databases.DB.Find(&record, id)
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).SendString("Couldn't find database record")
		}

		return c.Status(fiber.StatusOK).JSON(record)
	}
}

// create new database record
func CreateDatabase() func(c *fiber.Ctx) error {

	var record models.Database

	return func(c *fiber.Ctx) error {

		if err := c.BodyParser(&record); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
		}

		if result := databases.DB.Create(&record); result.RowsAffected == 0 {
			return c.Status(fiber.StatusConflict).SendString("Database already exists")
		}

		return c.Status(fiber.StatusCreated).JSON(record)
	}
}

// update an existing database record
func UpdateDatabase() func(c *fiber.Ctx) error {

	var record models.Database

	return func(c *fiber.Ctx) error {

		if err := c.BodyParser(record); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
		}

		id := c.Params("id")
		result := databases.DB.Where("id = ?", id).Updates(&record)
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).SendString("Couldn't find database record")
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

// delete a database record
func DeleteDatabase() func(c *fiber.Ctx) error {

	var dbRecord models.Database

	return func(c *fiber.Ctx) error {

		id := c.Params("id")
		result := databases.DB.Delete(&dbRecord, id)
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).SendString("Couldn't find database record")
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
