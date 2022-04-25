package handler

import (
	"errors"
	"testing"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type albumRepositoryMock struct {
	mock.Mock
}

func (mock *albumRepositoryMock) Create(album entity.Album) (entity.Album, error) {
	args := mock.Called(album)

	return album, args.Error(1)
}

func (mock *albumRepositoryMock) GetAll() ([]entity.Album, error) {
	args := mock.Called()

	return args.Get(0).([]entity.Album), args.Error(1)
}

func TestCreateAlbumSuccess(t *testing.T) {
	albumRepository := new(albumRepositoryMock)
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.00,
		Title:  "ANY_TITLE",
	}
	expectedResponse := entity.Response{
		StatusCode: 201,
		Message:    "Created",
		Data:       expectedAlbum,
	}
	albumHandler := AlbumHandlerREST{
		Repository: albumRepository,
	}

	albumRepository.On("Create", expectedAlbum).Return(expectedAlbum, nil)

	actualResponse := albumHandler.CreateAlbum(entity.CreateAlbumRequest{
		Artist: expectedAlbum.Artist,
		Price:  expectedAlbum.Price,
		Title:  expectedAlbum.Title,
	})

	assert.Equal(t, actualResponse, expectedResponse, "should return the success response")
}

func TestCreateAlbumError(t *testing.T) {
	albumRepository := new(albumRepositoryMock)
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.00,
		Title:  "ANY_TITLE",
	}
	expectedResponse := entity.Response{
		StatusCode: 500,
		Message:    "Internal Server Error",
		Data:       nil,
	}
	albumHandler := AlbumHandlerREST{
		Repository: albumRepository,
	}

	albumRepository.On("Create", expectedAlbum).Return(nil, errors.New("Create Error"))

	actualResponse := albumHandler.CreateAlbum(entity.CreateAlbumRequest{
		Artist: expectedAlbum.Artist,
		Price:  expectedAlbum.Price,
		Title:  expectedAlbum.Title,
	})

	assert.Equal(t, actualResponse, expectedResponse, "should return the error response")
}

func TestGetAlbumsSuccess(t *testing.T) {
	albumRepository := new(albumRepositoryMock)
	expectedAlbums := []entity.Album{
		{
			Artist: "ANY_ARTIST",
			Price:  100.00,
			Title:  "ANY_TITLE",
		},
	}
	expectedResponse := entity.Response{
		StatusCode: 200,
		Message:    "OK",
		Data:       expectedAlbums,
	}
	albumHandler := AlbumHandlerREST{
		Repository: albumRepository,
	}

	albumRepository.On("GetAll").Return(expectedAlbums, nil)

	actualResponse := albumHandler.GetAlbums()

	assert.Equal(t, actualResponse, expectedResponse, "should return the success response")
}

func TestGetAlbumsError(t *testing.T) {
	albumRepository := new(albumRepositoryMock)
	expectedResponse := entity.Response{
		StatusCode: 500,
		Message:    "Internal Server Error",
		Data:       nil,
	}
	albumHandler := AlbumHandlerREST{
		Repository: albumRepository,
	}

	albumRepository.On("GetAll").Return([]entity.Album{}, errors.New("Get Error"))

	actualResponse := albumHandler.GetAlbums()

	assert.Equal(t, actualResponse, expectedResponse, "should return the error response")
}
