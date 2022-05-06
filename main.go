package main

import (
	"github.com/ae-lexs/ae_albums_api/client"
	"github.com/ae-lexs/ae_albums_api/config"
	"github.com/ae-lexs/ae_albums_api/handler"
	"github.com/ae-lexs/ae_albums_api/repository"
	"github.com/ae-lexs/ae_albums_api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := config.Get()
	postgresClient := client.GetPostgres(
		config.DBName,
		config.DBHost,
		config.DBPassword,
		config.DBUser,
		config.DBPort,
	)
	albumRepository := repository.AlbumPostgres{
		Client: postgresClient,
	}
	albumHandlerREST := handler.AlbumREST{
		Repository: &albumRepository,
	}
	albumRoutes := route.AlbumRoute{
		Handler: &albumHandlerREST,
	}

	router.GET("/albums", albumRoutes.GetAlbums)
	router.POST("/albums", albumRoutes.CreateAlbum)

	router.Run("localhost:8080")
}
