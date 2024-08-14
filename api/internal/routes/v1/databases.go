package v1

import (
	"github.com/christian-nickerson/pangolin/api/internal/engines/databases"
	"github.com/christian-nickerson/pangolin/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

func AddDatabaseRoutes(route fiber.Router) {
	group := route.Group("/databases")

	group.Get("", GetDatabases)
	group.Get("/:databaseId", GetDatabase)
	group.Post("", CreateDatabase)
	group.Patch("/:databaseId", UpdateDatabase)
	group.Delete("/:databaseId", DeleteDatabase)
}

// return a list of all database records
func GetDatabases(c *fiber.Ctx) error {
	var dbRecords []models.Database

	databases.Client.Find(&dbRecords)
	return c.Status(200).JSON(dbRecords)
}

// return a specific database record
func GetDatabase(c *fiber.Ctx) error {
	id := c.Params("databaseId")
	var dbRecords models.Database

	result := databases.Client.Find(&dbRecords, id)
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

	databases.Client.Create(&dbRecord)
	return c.SendStatus(201)
}

// update an existing database record
func UpdateDatabase(c *fiber.Ctx) error {
	dbRecord := new(models.Database)
	id := c.Params("databaseId")

	if err := c.BodyParser(dbRecord); err != nil {
		return c.Status(422).SendString(err.Error())
	}

	result := databases.Client.Where("ID = ?", id).Updates(&dbRecord)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Couldn't find database record")
	}

	return c.SendStatus(200)
}

// delete a database record
func DeleteDatabase(c *fiber.Ctx) error {
	id := c.Params("databaseId")
	var dbRecord models.Database

	result := databases.Client.Delete(&dbRecord, id)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Couldn't find database record")
	}

	return c.SendStatus(200)
}
