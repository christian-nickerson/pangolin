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
	"github.com/stretchr/testify/suite"
)

type DatabaseEndpointsSuite struct {
	suite.Suite
	app      *fiber.App
	nRecords int
	pageSize int
	pages    int64
}

// set up app & DB
func (s *DatabaseEndpointsSuite) SetupTest() {
	// set up values
	s.nRecords = 13
	s.pageSize = 5
	s.pages = int64(math.Ceil(float64(s.nRecords) / float64(s.pageSize)))

	// set up app & database
	s.app = fiber.New()
	AddDatabaseRoutes(s.app)
	databases.Connect(&configs.DatabaseConfig{Type: "sqlite", DbName: "test"})
}

// build default records
func (s *DatabaseEndpointsSuite) BeforeTest(suiteName, testName string) {
	for i := range s.nRecords {
		name := fmt.Sprintf("test%v", i)
		record := models.Database{Name: name}
		databases.DB.Create(&record)
	}
}

// tear down all records after test
func (s *DatabaseEndpointsSuite) AfterTest(suiteName, testName string) {
	databases.DB.Where("id IS NOT NULL").Delete(&models.Database{})
}

// shutdown app & DB
func (s *DatabaseEndpointsSuite) TearDownTest() {
	s.app.Shutdown()
	databases.DB.Migrator().DropTable(&models.Database{})
}

// Test get databases paginates across all records correctly
func (s *DatabaseEndpointsSuite) TestGetDatabases() {
	var result models.DatabaseResponse
	var records []models.Database
	var continuationToken string
	var endPagination bool = false

	for endPagination != true {
		// build path
		path := fmt.Sprintf("/databases?pageSize=%v", s.pageSize)
		if continuationToken != "" {
			path = path + fmt.Sprintf("&continuationToken=%v", continuationToken)
		}

		// get body response
		request := httptest.NewRequest("GET", path, nil)
		response, _ := s.app.Test(request)
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		json.Unmarshal(body, &result)

		// test per request assertions
		s.Assert().Equal(200, response.StatusCode)
		s.Assert().Equal(int64(s.nRecords), result.TotalRecords)
		s.Assert().Equal(s.pages, result.TotalPages)

		// set up loop over
		continuationToken = result.ContinuationToken
		records = append(records, result.Data...)
		if len(result.Data) < s.pageSize {
			endPagination = true
		}
	}

	s.Assert().Equal(s.nRecords, len(records))
}

// Test get databases returns 204 when there are no records
func (s *DatabaseEndpointsSuite) TestGetDatabasesNoRecords() {
	// delete records
	databases.DB.Where("id IS NOT NULL").Delete(&models.Database{})

	// test request
	path := fmt.Sprintf("/databases?pageSize=%v", s.pageSize)
	request := httptest.NewRequest("GET", path, nil)
	response, _ := s.app.Test(request)

	s.Assert().Equal(204, response.StatusCode)
}

// Test get database returns correctly
func (s *DatabaseEndpointsSuite) TestGetDatabase() {
	var result models.Database

	request := httptest.NewRequest("GET", fmt.Sprintf("/databases/%v", 1), nil)
	response, _ := s.app.Test(request)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &result)

	s.Assert().Equal(200, response.StatusCode)
	s.Assert().Equal(result.Name, "test0")
}

// Test get database returns 404
func (s *DatabaseEndpointsSuite) TestGetDatabaseNotFound() {
	request := httptest.NewRequest("GET", fmt.Sprintf("/databases/%v", 100), nil)
	response, _ := s.app.Test(request)

	s.Assert().Equal(404, response.StatusCode)
}

func TestDatabaseEndpointsSuite(t *testing.T) {
	suite.Run(t, new(DatabaseEndpointsSuite))
}
