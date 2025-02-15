package v1

import "github.com/gofiber/fiber/v2"

func AddCollectionRoutes(route fiber.Router) {
	group := route.Group("databases/:dbId/collections")

	group.Get("", GetCollections)
	group.Get("/:id", GetCollection)
	group.Post("", CreateCollection)
	group.Patch("/:id", UpdateCollection)
	group.Delete("/:id", DeleteCollection)
	group.Post("/:id/search", SearchCollection)
}

func GetCollections(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func GetCollection(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func CreateCollection(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func UpdateCollection(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func DeleteCollection(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func SearchCollection(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
