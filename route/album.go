package route

import (
	"io/ioutil"
	"log"
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

func (route *album) Create(context *gin.Context) {
	requestBody, err := ioutil.ReadAll(context.Request.Body)

	if err != nil {
		log.Printf("albumRoute Create, Error with Request Body: %v", err)
		context.IndentedJSON(http.StatusInternalServerError, entity.Response{
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
	}

	response := route.handler.Create(requestBody)

	context.IndentedJSON(response.StatusCode, response)
}

func (route *album) Get(context *gin.Context) {
	response := route.handler.Get(context.Param("id"))

	context.IndentedJSON(response.StatusCode, response)
}
