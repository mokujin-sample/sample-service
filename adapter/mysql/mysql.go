package db

import (
	"sample-service/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(configuration *config.Database) (*gorm.DB, error) {
	dsn := configuration.User + ":" + configuration.Password + "@tcp(" + configuration.Host + ")/" + configuration.DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
