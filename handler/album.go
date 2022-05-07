package handler

import (
	"net/http"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/repository"
)

const (
	createdMessage             = "Created"
	internalServerErrorMessage = "Internal Server Error"
	okMessage                  = "OK"
)

type Album interface {
	Create(entity.CreateAlbumRequest) entity.Response
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

func (handler *albumREST) Create(receivedAlbum entity.CreateAlbumRequest) entity.Response {
	createdAlbum, err := handler.repository.Create(
		receivedAlbum.Artist,
		receivedAlbum.Price,
		receivedAlbum.Title,
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
