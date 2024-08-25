package v1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/christian-nickerson/pangolin/api/internal/engines/databases"
	"github.com/christian-nickerson/pangolin/api/internal/models"
)

// add database routes to a fiber app
func AddDatabaseRoutes(route fiber.Router) {
	group := route.Group("/databases")

	group.Get("", models.ValidatePagination, GetDatabases)
	group.Get("/:id", GetDatabase)
	group.Post("", models.ValidateDatabase, CreateDatabase)
	group.Patch("/:id", models.ValidateDatabase, UpdateDatabase)
	group.Delete("/:id", DeleteDatabase)
}

// return paginated database records
func GetDatabases(c *fiber.Ctx) error {
	var prq models.PaginationRequest
	var prp models.PaginationResponse
	var dbRecords []models.Database

	if err := c.QueryParser(&prq); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if result := databases.DB.Scopes(
		databases.Paginate(dbRecords, &prq, &prp, databases.DB),
	).Find(&dbRecords); result.RowsAffected == 0 {
		return c.Status(fiber.StatusNoContent).SendString("No records found")
	}

	dbRecords, prp.ContinuationToken = databases.PaginatedResponse(dbRecords, prq.PageSize)
	return c.Status(fiber.StatusOK).JSON(
		models.DatabaseResponse{
			Data:               dbRecords,
			PaginationResponse: prp,
		},
	)
}

// return a specific database record
func GetDatabase(c *fiber.Ctx) error {
	id := c.Params("id")
	var dbRecords models.Database

	// TODO: add where statement to correct IDs
	result := databases.DB.Find(&dbRecords, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Couldn't find database record")
	}

	return c.Status(fiber.StatusOK).JSON(dbRecords)
}

// create new database record
func CreateDatabase(c *fiber.Ctx) error {
	dbRecord := new(models.Database)

	if err := c.BodyParser(dbRecord); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	databases.DB.Create(&dbRecord)
	return c.SendStatus(fiber.StatusCreated)
}

// update an existing database record
func UpdateDatabase(c *fiber.Ctx) error {
	dbRecord := new(models.Database)
	id := c.Params("id")

	if err := c.BodyParser(dbRecord); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	result := databases.DB.Where("id = ?", id).Updates(&dbRecord)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Couldn't find database record")
	}

	return c.SendStatus(fiber.StatusOK)
}

// delete a database record
func DeleteDatabase(c *fiber.Ctx) error {
	id := c.Params("id")
	var dbRecord models.Database

	result := databases.DB.Delete(&dbRecord, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Couldn't find database record")
	}

	return c.SendStatus(fiber.StatusOK)
}
