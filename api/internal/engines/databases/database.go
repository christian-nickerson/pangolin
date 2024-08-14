package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/christian-nickerson/pangolin/api/internal/configs"
	"github.com/christian-nickerson/pangolin/api/internal/models"
)

var Client *gorm.DB

func Connect(config *configs.DatabseConfig) {
	var err error

	Client, err = gorm.Open(connector(config), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	Client.AutoMigrate(&models.Database{})
}

// build connection dialect object from config
func connector(config *configs.DatabseConfig) gorm.Dialector {
	var conn gorm.Dialector

	switch config.Type {
	case "postgres":
		conn = postgresConnector(config)
	case "sqlite":
		conn = sqliteConnector(config)
	default:
		log.Fatalf("Database %v is unsupported", config.Type)
	}

	return conn
}

// set postgres connection string and return conn object
func postgresConnector(config *configs.DatabseConfig) gorm.Dialector {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.DbName,
		config.Port,
	)
	return postgres.Open(dsn)
}

// set sqlite connetion string and return conn object
func sqliteConnector(config *configs.DatabseConfig) gorm.Dialector {
	return sqlite.Open(config.DbName)
}
