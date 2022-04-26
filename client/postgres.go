package client

import (
	"fmt"

	"github.com/ae-lexs/ae_albums_api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpPostgres(dbName, dbHost, dbPassword, dbUser, dbPort string) *gorm.DB {
	postgresConfig := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)
	postgresClient, err := gorm.Open(
		postgres.Open(
			postgresConfig,
		),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed to connect to the database")
	}

	postgresClient.AutoMigrate(&entity.Album{})

	return postgresClient
}
