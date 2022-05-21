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

func (mock *albumRepositoryMock) GetByID(albumID string) (entity.Album, error) {
	args := mock.Called(albumID)

	return args.Get(0).(entity.Album), args.Error(1)
}

func TestAlbumHandlerCreate(t *testing.T) {
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.00,
		Title:  "ANY_TITLE",
	}
	testCases := []struct {
		expectedResponse entity.Response
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

			assert.Equal(t, actualResponse, testCase.expectedResponse)
		})
	}
}

func TestAlbumHandlerGet(t *testing.T) {
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.00,
		Title:  "ANY_TITLE",
	}
	expectedAlbums := []entity.Album{expectedAlbum}
	testCases := []struct {
		albumID          string
		expectedResponse entity.Response
		mockAllResponse  []entity.Album
		mockOneResponse  entity.Album
		mockAllError     error
		mockOneError     error
		name             string
	}{
		{
			albumID: "",
			expectedResponse: entity.Response{
				StatusCode: 200,
				Message:    "OK",
				Data:       expectedAlbums,
			},
			mockAllResponse: expectedAlbums,
			mockAllError:    nil,
			name:            "OK Response Get All",
		},
		{
			albumID: "",
			expectedResponse: entity.Response{
				StatusCode: 500,
				Message:    "Internal Server Error",
				Data:       nil,
			},
			mockAllResponse: []entity.Album{},
			mockAllError:    errors.New("GetError"),
			name:            "Internal Server Error Response Get All",
		},
		{
			albumID: "1",
			expectedResponse: entity.Response{
				StatusCode: 200,
				Message:    "OK",
				Data:       expectedAlbum,
			},
			mockOneResponse: expectedAlbum,
			mockOneError:    nil,
			name:            "OK Response Get Specific",
		},
		{
			albumID: "1",
			expectedResponse: entity.Response{
				StatusCode: 404,
				Message:    "Not Found",
				Data:       nil,
			},
			mockOneResponse: entity.Album{},
			mockOneError:    errors.New("AlbumNotFoundError"),
			name:            "Not Found Response Get Specific",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			albumRepository := new(albumRepositoryMock)
			albumHandler := handler.NewAlbumREST(albumRepository)

			albumRepository.On("GetAll").Return(testCase.mockAllResponse, testCase.mockAllError)
			albumRepository.On("GetByID", testCase.albumID).Return(
				testCase.mockOneResponse,
				testCase.mockOneError,
			)

			actualResponse := albumHandler.Get(testCase.albumID)

			assert.Equal(t, actualResponse, testCase.expectedResponse)
		})
	}
}
