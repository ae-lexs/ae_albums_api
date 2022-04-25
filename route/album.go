package route

import (
	"net/http"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/handler"
	"github.com/gin-gonic/gin"
)

type AlbumRoute struct {
	Handler handler.AlbumHandler
}

func (route *AlbumRoute) CreateAlbum(c *gin.Context) {
	var receivedAlbum entity.CreateAlbumRequest

	if err := c.BindJSON(&receivedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
	}

	response := route.Handler.CreateAlbum(receivedAlbum)

	c.IndentedJSON(response.StatusCode, response)
}

func (route *AlbumRoute) GetAlbums(c *gin.Context) {
	response := route.Handler.GetAlbums()

	c.IndentedJSON(response.StatusCode, response)
}

// func getAlbumById(c *gin.Context) {
// 	receivedId := c.Param("id")

// 	for _, album := range albums {
// 		if album.ID == receivedId {
// 			c.IndentedJSON(http.StatusOK, album)

// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album Not Found"})
// }
