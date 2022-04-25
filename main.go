package main

import (
	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/handler"
	"github.com/ae-lexs/ae_albums_api/repository"
	"github.com/ae-lexs/ae_albums_api/route"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpDB() *gorm.DB {
	postgresClient, err := gorm.Open(
		postgres.Open(
			"host=localhost user=album_user password=album2022 dbname=album_db port=5433 sslmode=disable",
		),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed to connect to the database")
	}

	postgresClient.AutoMigrate(&entity.Album{})

	return postgresClient
}

func main() {
	router := gin.Default()
	postgresClient := setUpDB()
	albumRepository := repository.PostgresAlbumRepository{
		Client: postgresClient,
	}
	albumHandlerREST := handler.AlbumHandlerREST{
		Repository: &albumRepository,
	}
	albumRoutes := route.AlbumRoute{
		Handler: &albumHandlerREST,
	}

	router.GET("/albums", albumRoutes.GetAlbums)
	router.POST("/albums", albumRoutes.CreateAlbum)

	router.Run("localhost:8080")
}
