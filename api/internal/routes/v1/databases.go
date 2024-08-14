package v1

import "github.com/gofiber/fiber/v2"

func AddDatabaseRoutes(route fiber.Router) {
	group := route.Group("/databases")

	group.Get("", GetDatabases)
	group.Get("/:databaseId", GetDatabase)
	group.Post("", CreateDatabase)
	group.Patch("/:databaseId", UpdateDatabase)
	group.Delete("/:databaseId", DeleteDatabase)
}

func GetDatabases(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func GetDatabase(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func CreateDatabase(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func UpdateDatabase(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func DeleteDatabase(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
