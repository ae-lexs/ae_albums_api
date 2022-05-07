package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/repository"
)

const (
	badRequestMessage          = "Bad Request"
	createdMessage             = "Created"
	internalServerErrorMessage = "Internal Server Error"
	okMessage                  = "OK"
)

type Album interface {
	Create([]byte) entity.Response
	Get() entity.Response
}

type albumREST struct {
	repository repository.Album
}

func NewAlbumREST(repository repository.Album) albumREST {
	return albumREST{
		repository: repository,
	}
}

func (handler *albumREST) Create(requestBody []byte) entity.Response {
	var albumData entity.CreateAlbumRequest

	if err := json.Unmarshal(requestBody, &albumData); err != nil {
		log.Printf("albumREST Create, Error Parsing JSON: %v", err)

		return entity.Response{
			StatusCode: http.StatusBadRequest,
			Message:    badRequestMessage,
			Data:       nil,
		}
	}

	createdAlbum, err := handler.repository.Create(
		albumData.Artist,
		albumData.Price,
		albumData.Title,
	)

	if err != nil {
		return entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    internalServerErrorMessage,
			Data:       nil,
		}
	}

	return entity.Response{
		StatusCode: http.StatusCreated,
		Message:    createdMessage,
		Data:       createdAlbum,
	}
}

func (handler *albumREST) Get() entity.Response {
	foundAlbums, err := handler.repository.GetAll()

	if err != nil {
		return entity.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    internalServerErrorMessage,
			Data:       nil,
		}
	}

	return entity.Response{
		StatusCode: http.StatusOK,
		Message:    okMessage,
		Data:       foundAlbums,
	}
}
