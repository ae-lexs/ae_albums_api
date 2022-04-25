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

type AlbumHandler interface {
	CreateAlbum(entity.CreateAlbumRequest) entity.Response
	GetAlbums() entity.Response
}

type AlbumHandlerREST struct {
	Repository repository.AlbumRepository
}

func (handler *AlbumHandlerREST) CreateAlbum(receivedAlbum entity.CreateAlbumRequest) entity.Response {
	newAlbum := entity.Album{
		Artist: receivedAlbum.Artist,
		Price:  receivedAlbum.Price,
		Title:  receivedAlbum.Title,
	}

	createdAlbum, err := handler.Repository.Create(newAlbum)

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

func (handler *AlbumHandlerREST) GetAlbums() entity.Response {
	foundAlbums, err := handler.Repository.GetAll()

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
