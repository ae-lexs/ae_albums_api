package route

import (
	"net/http"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/repository"
	"github.com/gin-gonic/gin"
)

type AlbumRoute struct {
	Repository repository.AlbumRepository
}

func (a *AlbumRoute) GetAlbums(c *gin.Context) {
	albums, err := a.Repository.GetAll()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func (a *AlbumRoute) CreateAlbum(c *gin.Context) {
	var newAlbum entity.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	if _, err := a.Repository.Create(newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
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
