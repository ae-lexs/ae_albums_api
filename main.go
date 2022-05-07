package main

import (
	"github.com/ae-lexs/ae_albums_api/client"
	"github.com/ae-lexs/ae_albums_api/config"
	"github.com/ae-lexs/ae_albums_api/handler"
	"github.com/ae-lexs/ae_albums_api/model"
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
	albumModel := model.NewAlbumPostgres(postgresClient)
	albumRepository := repository.NewAlbum(&albumModel)
	albumHandlerREST := handler.NewAlbumREST(&albumRepository)
	albumRoutes := route.NewAlbum(&albumHandlerREST)

	router.GET("/albums", albumRoutes.Get)
	router.POST("/albums", albumRoutes.Create)

	router.Run("localhost:8080")
}
