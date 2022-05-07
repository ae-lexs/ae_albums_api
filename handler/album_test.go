package handler_test

import (
	"errors"
	"testing"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type albumRepositoryMock struct {
	mock.Mock
}

func (mock *albumRepositoryMock) Create(artist string, price float64, title string) (entity.Album, error) {
	args := mock.Called(artist, price, title)

	return args.Get(0).(entity.Album), args.Error(1)
}

func (mock *albumRepositoryMock) GetAll() ([]entity.Album, error) {
	args := mock.Called()

	return args.Get(0).([]entity.Album), args.Error(1)
}

func TestCreate(t *testing.T) {
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.00,
		Title:  "ANY_TITLE",
	}
	testCases := []struct {
		expectedResponse entity.Response
		receivedAlbum    entity.CreateAlbumRequest
		message          string
		mockResponse     entity.Album
		mockError        error
		name             string
	}{
		{
			expectedResponse: entity.Response{
				StatusCode: 201,
				Message:    "Created",
				Data:       expectedAlbum,
			},
			receivedAlbum: entity.CreateAlbumRequest{
				Artist: expectedAlbum.Artist,
				Price:  expectedAlbum.Price,
				Title:  expectedAlbum.Title,
			},
			message:      "should return the OK Response",
			mockResponse: expectedAlbum,
			mockError:    nil,
			name:         "OK Response",
		},
		{
			expectedResponse: entity.Response{
				StatusCode: 500,
				Message:    "Internal Server Error",
				Data:       nil,
			},
			receivedAlbum: entity.CreateAlbumRequest{
				Artist: expectedAlbum.Artist,
				Price:  expectedAlbum.Price,
				Title:  expectedAlbum.Title,
			},
			message:      "should return the Internal Server Error Response",
			mockResponse: expectedAlbum,
			mockError:    errors.New("Repository Create Error"),
			name:         "Internal Server Error Response",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			albumRepository := new(albumRepositoryMock)
			albumHandler := handler.NewAlbumREST(albumRepository)

			albumRepository.On(
				"Create",
				testCase.receivedAlbum.Artist,
				testCase.receivedAlbum.Price,
				testCase.receivedAlbum.Title,
			).Return(
				testCase.mockResponse,
				testCase.mockError,
			)

			actualResponse := albumHandler.Create(testCase.receivedAlbum)

			assert.Equal(t, actualResponse, testCase.expectedResponse, testCase.message)
		})
	}
}

func TestGet(t *testing.T) {
	expectedAlbums := []entity.Album{
		{
			Artist: "ANY_ARTIST",
			Price:  100.00,
			Title:  "ANY_TITLE",
		},
	}
	testCases := []struct {
		expectedResponse entity.Response
		message          string
		mockResponse     []entity.Album
		mockError        error
		name             string
	}{
		{
			expectedResponse: entity.Response{
				StatusCode: 200,
				Message:    "OK",
				Data:       expectedAlbums,
			},
			message:      "should return the OK Response",
			mockResponse: expectedAlbums,
			mockError:    nil,
			name:         "OK Response",
		},
		{
			expectedResponse: entity.Response{
				StatusCode: 500,
				Message:    "Internal Server Error",
				Data:       nil,
			},
			message:      "should return the Internal Server Error Response",
			mockResponse: []entity.Album{},
			mockError:    errors.New("Get Error"),
			name:         "Internal Server Error Response",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			albumRepository := new(albumRepositoryMock)
			albumHandler := handler.NewAlbumREST(albumRepository)

			albumRepository.On("GetAll").Return(testCase.mockResponse, testCase.mockError)

			actualResponse := albumHandler.Get()

			assert.Equal(t, actualResponse, testCase.expectedResponse, testCase.message)
		})
	}
}
