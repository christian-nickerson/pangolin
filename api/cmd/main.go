package main

import (
	"log"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"

	"github.com/christian-nickerson/pangolin/api/internal/configs"
	"github.com/christian-nickerson/pangolin/api/internal/routes"
)

func main() {
	// load settings
	settings, err := configs.LoadSettings("settings")
	if err != nil {
		log.Fatal("Loading settings failed:", err)
	}

	// set up fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// add routes
	app.Use(healthcheck.New(routes.HealthCheckConfig))

	// start serving
	app.Listen(":" + strconv.Itoa(settings.Server.API.Port))
}
