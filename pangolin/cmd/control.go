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

	"github.com/christian-nickerson/pangolin/pangolin/internal/configs"
	"github.com/christian-nickerson/pangolin/pangolin/internal/logging"
	"github.com/christian-nickerson/pangolin/pangolin/internal/routes/health"
)

// Build & run control plane
func startService(settings *configs.Settings) *fiber.App {

	// configure fiber app
	app := fiber.New(fiber.Config{
		AppName:               "Pangolin",
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	// middleware and routes
	app.Use(requestid.New())
	app.Use(logger.New(logging.LoggingConfig))
	app.Use(healthcheck.New(health.HealthCheckConfig))

	// start serving in new goroutine
	go func() {
		address := fmt.Sprintf("127.0.0.1:%v", settings.Server.API.Port)
		if err := app.Listen(address); err != nil {
			log.Fatal(err.Error())
		}
	}()

	return app
}

// Handle IO to service & run
func main() {

	// handle interruptions
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Load dependent objects
	settings, err := configs.Load("settings.toml")
	if err != nil {
		log.Fatal(err.Error())
	}

	// start service and wait for signal
	app := startService(&settings)
	log.Printf("Started serving on http://127.0.0.1:%v\n", settings.Server.API.Port)
	<-ctx.Done()

	log.Println("Starting shutting down...")

	// shutdown and close connections
	if err := app.Shutdown(); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Pangolin successfully shutdown.")
}
