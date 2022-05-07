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

func TestAlbumHandlerCreate(t *testing.T) {
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.00,
		Title:  "ANY_TITLE",
	}
	testCases := []struct {
		expectedResponse entity.Response
		message          string
		mockError        error
		mockInputArtist  string
		mockInputPrice   float64
		mockInputTitle   string
		mockResponse     entity.Album
		name             string
		requestData      []byte
	}{
		{
			expectedResponse: entity.Response{
				StatusCode: 201,
				Message:    "Created",
				Data:       expectedAlbum,
			},
			message:         "should return the OK Response",
			mockError:       nil,
			mockInputArtist: expectedAlbum.Artist,
			mockInputPrice:  expectedAlbum.Price,
			mockInputTitle:  expectedAlbum.Title,
			mockResponse:    expectedAlbum,
			name:            "OK Response",
			requestData: []byte(`{
				"artist": "ANY_ARTIST",
				"price":  "100.00",
				"title":  "ANY_TITLE"
			}`),
		},
		{
			expectedResponse: entity.Response{
				StatusCode: 400,
				Message:    "Bad Request",
				Data:       nil,
			},
			message:         "should return the Bad Request Response",
			mockError:       nil,
			mockInputArtist: "",
			mockInputPrice:  0,
			mockInputTitle:  "",
			mockResponse:    entity.Album{},
			name:            "Bad Request Response",
			requestData: []byte(`{
				"artist": "ANY_ARTIST",
				"price":  "100.00",
				"title":  "ANY_TITLE",
			}`),
		},
		{
			expectedResponse: entity.Response{
				StatusCode: 500,
				Message:    "Internal Server Error",
				Data:       nil,
			},
			message:         "should return the Internal Server Error Response",
			mockError:       errors.New("Repository Create Error"),
			mockInputArtist: expectedAlbum.Artist,
			mockInputPrice:  expectedAlbum.Price,
			mockInputTitle:  expectedAlbum.Title,
			mockResponse:    expectedAlbum,
			name:            "Internal Server Error Response",
			requestData: []byte(`{
				"artist": "ANY_ARTIST",
				"price":  "100.00",
				"title":  "ANY_TITLE"
			}`),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			albumRepository := new(albumRepositoryMock)
			albumHandler := handler.NewAlbumREST(albumRepository)

			albumRepository.On(
				"Create",
				testCase.mockInputArtist,
				testCase.mockInputPrice,
				testCase.mockInputTitle,
			).Return(
				testCase.mockResponse,
				testCase.mockError,
			)

			actualResponse := albumHandler.Create(testCase.requestData)

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
