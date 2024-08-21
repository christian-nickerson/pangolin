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

// return paginated databse records
func GetDatabases(c *fiber.Ctx) error {
	var pagination models.PaginationRequest
	var dbRecords []models.Database

	c.QueryParser(&pagination)
	databases.DB.Scopes(databases.Paginate(&pagination)).Find(&dbRecords)
	nextToken := databases.GetContinuationToken(dbRecords)

	response := models.DatabaseResponse{
		Databases: dbRecords[:len(dbRecords)-1],
		PaginationResponse: models.PaginationResponse{
			ContinuationToken: nextToken,
		},
	}

	return c.Status(200).JSON(response)
}

// return a specific database record
func GetDatabase(c *fiber.Ctx) error {
	id := c.Params("id")
	var dbRecords models.Database

	// TODO: add where statement to correct IDs
	result := databases.DB.Find(&dbRecords, id)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Couldn't find database record")
	}

	return c.Status(200).JSON(dbRecords)
}

// create new database record
func CreateDatabase(c *fiber.Ctx) error {
	dbRecord := new(models.Database)

	if err := c.BodyParser(dbRecord); err != nil {
		return c.Status(422).SendString(err.Error())
	}

	databases.DB.Create(&dbRecord)
	return c.SendStatus(201)
}

// update an existing database record
func UpdateDatabase(c *fiber.Ctx) error {
	dbRecord := new(models.Database)
	id := c.Params("id")

	if err := c.BodyParser(dbRecord); err != nil {
		return c.Status(422).SendString(err.Error())
	}

	result := databases.DB.Where("id = ?", id).Updates(&dbRecord)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Couldn't find database record")
	}

	return c.SendStatus(200)
}

// delete a database record
func DeleteDatabase(c *fiber.Ctx) error {
	id := c.Params("id")
	var dbRecord models.Database

	result := databases.DB.Delete(&dbRecord, id)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Couldn't find database record")
	}

	return c.SendStatus(200)
}
