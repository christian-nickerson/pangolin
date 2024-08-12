package main

import (
	"log"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/christian-nickerson/pangolin/api/internal/configs"
	"github.com/christian-nickerson/pangolin/api/internal/logging"
	"github.com/christian-nickerson/pangolin/api/internal/routes"
)

func main() {
	// load settings
	settings, err := configs.Load("settings.toml")
	if err != nil {
		log.Fatal("Loading settings failed:", err)
	}

	// set up fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// add logging
	app.Use(requestid.New())
	app.Use(logger.New(logging.LoggingConfig))

	// add routes
	app.Use(healthcheck.New(routes.HealthCheckConfig))

	// start serving
	app.Listen(":" + strconv.Itoa(settings.Server.API.Port))
}
