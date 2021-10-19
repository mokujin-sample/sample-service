package db

import (
	"object-service/config"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// Connect to a database handle from a connection string.
func Connect(configuration *config.Database) (*gorm.DB, error) {
	dsn := "tcp://" + configuration.Host + ":" + configuration.Port + "?database=" + configuration.DB + "&read_timeout=10"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
