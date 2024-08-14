package v1

import "github.com/gofiber/fiber/v2"

func AddCollectionRoutes(route fiber.Router) {
	group := route.Group("/collections")

	group.Get(":databaseId/", GetCollections)
	group.Get(":databaseId/:collectionId", GetCollection)
	group.Post(":databaseId/", CreateCollection)
	group.Patch(":databaseId/:collectionId", UpdateCollection)
	group.Delete(":databaseId/:collectionId", DeleteCollection)
}

func GetCollections(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func GetCollection(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func CreateCollection(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func UpdateCollection(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func DeleteCollection(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
