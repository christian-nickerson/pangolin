package v1

import (
	"encoding/base64"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"

	"github.com/christian-nickerson/pangolin/api/internal/engines/databases"
	"github.com/christian-nickerson/pangolin/api/internal/models"
)

// add database routes to a fiber app
func AddDatabaseRoutes(route fiber.Router) {
	group := route.Group("/databases")

	group.Get("", models.ValidatePagination, GetDatabases)
	group.Get("/:databaseId", GetDatabase)
	group.Post("", CreateDatabase)
	group.Patch("/:databaseId", UpdateDatabase)
	group.Delete("/:databaseId", DeleteDatabase)
}

// return paginated databse records
func GetDatabases(c *fiber.Ctx) error {
	var pagination models.PaginationRequest
	var dbRecords []models.Database

	c.QueryParser(&pagination)

	// set up base query without cursor token
	query := databases.DB.Limit(pagination.PageSize + 1).Order(clause.OrderByColumn{
		Column: clause.Column{Name: "id"},
		Desc:   pagination.OrderDesc,
	})

	if pagination.ContinuationToken != "" {
		// base64 validation handled earlier by validator
		cursor, _ := base64.StdEncoding.DecodeString(pagination.ContinuationToken)
		query = query.Where("id <= ?", cursor)
	}

	query.Find(&dbRecords)

	n1 := dbRecords[len(dbRecords)-1]
	idByteString := []byte(strconv.FormatUint(uint64(n1.ID), 10))
	nextToken := base64.StdEncoding.EncodeToString(idByteString)

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
	id := c.Params("databaseId")
	var dbRecords models.Database

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
	id := c.Params("databaseId")

	if err := c.BodyParser(dbRecord); err != nil {
		return c.Status(422).SendString(err.Error())
	}

	result := databases.DB.Where("ID = ?", id).Updates(&dbRecord)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Couldn't find database record")
	}

	return c.SendStatus(200)
}

// delete a database record
func DeleteDatabase(c *fiber.Ctx) error {
	id := c.Params("databaseId")
	var dbRecord models.Database

	result := databases.DB.Delete(&dbRecord, id)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Couldn't find database record")
	}

	return c.SendStatus(200)
}
