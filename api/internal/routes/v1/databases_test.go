package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http/httptest"
	"testing"

	"github.com/christian-nickerson/pangolin/api/internal/configs"
	"github.com/christian-nickerson/pangolin/api/internal/engines/databases"
	"github.com/christian-nickerson/pangolin/api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var databaseSettings configs.DatabaseConfig = configs.DatabaseConfig{
	Type:     "sqlite",
	Host:     "test",
	Port:     0,
	DbName:   "test",
	Username: "test",
	Password: "test",
}

// Test get databases returns correct records
func TestGetDatabases(t *testing.T) {
	var nRecords int64 = 13
	var pageSize int = 5
	var pages int64 = int64(math.Ceil(float64(nRecords) / float64(pageSize)))

	// set up database
	databases.Connect(&databaseSettings)
	for i := range nRecords {
		name := fmt.Sprintf("test%v", i)
		record := models.Database{Name: name}
		databases.DB.Create(&record)
	}

	// setup app
	app := fiber.New()
	AddDatabaseRoutes(app)

	// test request
	path := fmt.Sprintf("/databases?pageSize=%v", pageSize)
	request := httptest.NewRequest("GET", path, nil)
	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)

	// test response body
	var result models.DatabaseResponse
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &result)

	assert.Equal(t, pageSize, len(result.Data))
	assert.Equal(t, nRecords, result.TotalRecords)
	assert.Equal(t, pages, result.TotalPages)
}

// Test get databases returns 204 when there are no records
// func TestGetDatabasesNoRecords(t *testing.T) {
// 	// setup app
// 	app := fiber.New()
// 	AddDatabaseRoutes(app)
// 	err := databases.Connect(&databaseSettings)
// 	assert.NoError(t, err, nil, "failed to connect to database")

// 	// test request
// 	path := fmt.Sprintf("/databases?pageSize=%v", 5)
// 	request := httptest.NewRequest("GET", path, nil)
// 	response, _ := app.Test(request)

// 	assert.Equal(t, 204, response.StatusCode)
// }
