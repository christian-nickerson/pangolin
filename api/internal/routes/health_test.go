package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/stretchr/testify/assert"
)

// Test health endpoint returns correctly
func TestHealthEndpoint(t *testing.T) {
	// setup app
	app := fiber.New()
	app.Use(healthcheck.New(HealthCheckConfig))

	// test request
	request := httptest.NewRequest("GET", "/health", nil)
	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
}

// Test health endpoint returns correctly
func TestReadyEndpoint(t *testing.T) {
	// setup app
	app := fiber.New()
	app.Use(healthcheck.New(HealthCheckConfig))

	// test request
	request := httptest.NewRequest("GET", "/ready", nil)
	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
}
