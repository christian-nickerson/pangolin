package v1

import "github.com/gofiber/fiber/v2"

// Add all V1 routes to a Fiber App
func AddV1Routes(router fiber.Router) {
	v1 := router.Group("/v1")

	AddDatabaseRoutes(v1)
	AddCollectionRoutes(v1)
	AddDocumentsRoutes(v1)
}
