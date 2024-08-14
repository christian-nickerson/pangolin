package main

import (
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/christian-nickerson/pangolin/api/internal/configs"
	"github.com/christian-nickerson/pangolin/api/internal/logging"
	"github.com/christian-nickerson/pangolin/api/internal/routes/health"
	v1 "github.com/christian-nickerson/pangolin/api/internal/routes/v1"
)

func main() {
	// load settings
	settings := configs.Load("settings.toml")

	// set up fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// add logging
	app.Use(requestid.New())
	app.Use(logger.New(logging.LoggingConfig))

	// add routes
	app.Use(healthcheck.New(health.HealthCheckConfig))
	v1.AddV1Routes(app)

	// start serving
	app.Listen(":" + strconv.Itoa(settings.Server.API.Port))
}
