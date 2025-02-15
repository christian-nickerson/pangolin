package databases

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/christian-nickerson/pangolin/pangolin/internal/configs"
	"github.com/christian-nickerson/pangolin/pangolin/internal/models"
)

var DB *gorm.DB

// connect to databases
func Connect(config *configs.DatabaseConfig) error {
	var err error

	connection, err := connector(config)
	if err != nil {
		return err
	}

	DB, err = gorm.Open(connection, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return fmt.Errorf("failure connecting to database, %v", err)
	}

	if err = DB.AutoMigrate(&models.Database{}); err != nil {
		return fmt.Errorf("failure creating Database schema, %v", err)
	}

	return err
}

func Close() error {
	instance, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failure returning DB instance, %v", err)
	}

	if err = instance.Close(); err != nil {
		return fmt.Errorf("failure closing connection, %v", err)
	}

	return nil
}

// build connection dialect object from config
func connector(config *configs.DatabaseConfig) (gorm.Dialector, error) {
	var conn gorm.Dialector
	var err error = nil

	switch config.Type {
	case "postgres":
		conn = postgresConnector(config)
	case "sqlite":
		conn = sqliteConnector(config)
	default:
		err = fmt.Errorf("database %v is unsupported", config.Type)
	}

	return conn, err
}

// set postgres connection string and return conn object
func postgresConnector(config *configs.DatabaseConfig) gorm.Dialector {
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
func sqliteConnector(config *configs.DatabaseConfig) gorm.Dialector {
	fileName := fmt.Sprintf("%v.db", config.DbName)
	return sqlite.Open(fileName)
}
