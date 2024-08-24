package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/christian-nickerson/pangolin/api/internal/configs"
	"github.com/christian-nickerson/pangolin/api/internal/engines/databases"
	"github.com/christian-nickerson/pangolin/api/internal/logging"
	"github.com/christian-nickerson/pangolin/api/internal/routes/health"
	v1 "github.com/christian-nickerson/pangolin/api/internal/routes/v1"
)

// Run the Fiber app
func startService(settings *configs.Settings) (*fiber.App, error) {
	// configure fiber app
	app := fiber.New(fiber.Config{
		AppName:               "Pangolin",
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: false,
	})

	// logging
	app.Use(requestid.New())
	app.Use(logger.New(logging.LoggingConfig))

	// routes
	app.Use(healthcheck.New(health.HealthCheckConfig))
	v1.AddV1Routes(app)

	// serve
	address := fmt.Sprintf("127.0.0.1:%v", settings.Server.API.Port)
	err := app.Listen(address)

	return app, err
}

func main() {
	// handle interruptions
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Load settings
	settings, err := configs.Load("settings.toml")
	if err != nil {
		log.Fatalf("Settings failure: %v", err)
	}

	// Connect to metadata DB
	if err := databases.Connect(&settings.Metadata.Database); err != nil {
		log.Fatalf("Metadata database failure: %v", err)
	}

	// start service
	app, err := startService(&settings)
	if err != nil {
		log.Fatalf("Service startup failure: %v\n", err)
	}

	log.Printf("Started serving on http://127.0.0.1:%v\n", settings.Server.API.Port)

	// enter graceful shutdown
	<-ctx.Done()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Service shutdown failure: %v\n", err)
	}

	if err := databases.Close(); err != nil {
		log.Fatalf("Metadata database closing failure: %v\n", err)
	}
}
