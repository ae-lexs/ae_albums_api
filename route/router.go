package route

import (
	"github.com/ae-lexs/ae_albums_api/handler"
	"github.com/ae-lexs/ae_albums_api/model"
	"github.com/ae-lexs/ae_albums_api/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGinRouter(postgresClient *gorm.DB) *gin.Engine {
	router := gin.Default()
	albumModel := model.NewAlbumPostgres(postgresClient)
	albumRepository := repository.NewAlbum(&albumModel)
	albumHandlerREST := handler.NewAlbumREST(&albumRepository)
	albumRoutes := NewAlbum(&albumHandlerREST)

	album := router.Group("/album")

	album.GET("/", albumRoutes.Get)
	album.GET("/:id", albumRoutes.Get)
	album.POST("/", albumRoutes.Create)

	return router
}
