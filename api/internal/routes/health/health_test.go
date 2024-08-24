package health

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/stretchr/testify/suite"
)

type HealthCheckSuite struct {
	suite.Suite
	app *fiber.App
}

// set up app
func (s *HealthCheckSuite) SetupTest() {
	s.app = fiber.New()
	s.app.Use(healthcheck.New(HealthCheckConfig))
}

// shutdown app
func (s *HealthCheckSuite) TearDownTest() {
	s.app.Shutdown()
}

// Test health endpoint returns correctly
func (s *HealthCheckSuite) TestHealthEndpoint() {
	request := httptest.NewRequest("GET", "/health", nil)
	response, _ := s.app.Test(request)

	s.Assert().Equal(200, response.StatusCode)
}

// Test health endpoint returns correctly
func (s *HealthCheckSuite) TestReadyEndpoint() {
	request := httptest.NewRequest("GET", "/ready", nil)
	response, _ := s.app.Test(request)

	s.Assert().Equal(200, response.StatusCode)
}

func TestHealthCheckSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckSuite))
}
