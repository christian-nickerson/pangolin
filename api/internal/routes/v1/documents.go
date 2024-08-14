package v1

import "github.com/gofiber/fiber/v2"

func AddDocumentsRoutes(route fiber.Router) {
	group := route.Group("/documents")

	group.Get(":databaseId/:collectionId", GetDocuments)
	group.Get(":databaseId/:collectionId/:documentId", GetDocument)
	group.Post(":databaseId/:collectionId", CreateDocument)
	group.Post(":databaseId/:collectionId/search", SearchDocument)
	group.Patch(":databaseId/:collectionId/:documentId", UpdateDocument)
	group.Delete(":databaseId/:collectionId/:documentId", DeleteDocument)
}

func GetDocuments(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func GetDocument(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func CreateDocument(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func UpdateDocument(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func DeleteDocument(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func SearchDocument(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
