package v1

import "github.com/gofiber/fiber/v2"

func AddDocumentsRoutes(route fiber.Router) {
	group := route.Group("databases/:dbId/collections/:colId/documents")

	group.Get("", GetDocuments)
	group.Get("/:id", GetDocument)
	group.Post(":/id", CreateDocument)
	group.Patch("/:id", UpdateDocument)
	group.Delete("/:id", DeleteDocument)
}

func GetDocuments(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func GetDocument(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func CreateDocument(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func UpdateDocument(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func DeleteDocument(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
