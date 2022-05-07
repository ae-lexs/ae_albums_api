package route

import (
	"net/http"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/handler"
	"github.com/gin-gonic/gin"
)

type album struct {
	handler handler.Album
}

func NewAlbum(handler handler.Album) album {
	return album{
		handler: handler,
	}
}

func (route *album) Create(c *gin.Context) {
	var receivedAlbum entity.CreateAlbumRequest

	if err := c.BindJSON(&receivedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
	}

	response := route.handler.Create(receivedAlbum)

	c.IndentedJSON(response.StatusCode, response)
}

func (route *album) Get(c *gin.Context) {
	response := route.handler.Get()

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
