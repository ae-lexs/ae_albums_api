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

func (mock *albumRepositoryMock) Create(album entity.Album) (entity.Album, error) {
	args := mock.Called(album)

	return album, args.Error(1)
}

func (mock *albumRepositoryMock) GetAll() ([]entity.Album, error) {
	args := mock.Called()

	return args.Get(0).([]entity.Album), args.Error(1)
}

func TestCreateAlbum(t *testing.T) {
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.00,
		Title:  "ANY_TITLE",
	}
	testCases := []struct {
		expectedResponse   entity.Response
		input              entity.CreateAlbumRequest
		message            string
		mockMethodInput    entity.Album
		mockMethodResponse entity.Album
		mockMethodError    error
		name               string
	}{
		{
			expectedResponse: entity.Response{
				StatusCode: 201,
				Message:    "Created",
				Data:       expectedAlbum,
			},
			input: entity.CreateAlbumRequest{
				Artist: expectedAlbum.Artist,
				Price:  expectedAlbum.Price,
				Title:  expectedAlbum.Title,
			},
			message:            "should return the OK Response",
			mockMethodInput:    expectedAlbum,
			mockMethodResponse: expectedAlbum,
			mockMethodError:    nil,
			name:               "OK Response",
		},
		{
			expectedResponse: entity.Response{
				StatusCode: 500,
				Message:    "Internal Server Error",
				Data:       nil,
			},
			input: entity.CreateAlbumRequest{
				Artist: expectedAlbum.Artist,
				Price:  expectedAlbum.Price,
				Title:  expectedAlbum.Title,
			},
			message:            "should return the Internal Server Error Response",
			mockMethodInput:    expectedAlbum,
			mockMethodResponse: expectedAlbum,
			mockMethodError:    errors.New("Create Error"),
			name:               "Internal Server Error Response",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			albumRepository := new(albumRepositoryMock)
			albumHandler := handler.NewAlbumREST(albumRepository)

			albumRepository.On("Create", testCase.mockMethodInput).Return(testCase.mockMethodResponse, testCase.mockMethodError)

			actualResponse := albumHandler.CreateAlbum(testCase.input)

			assert.Equal(t, actualResponse, testCase.expectedResponse, testCase.message)
		})
	}
}

func TestGetAlbums(t *testing.T) {
	expectedAlbums := []entity.Album{
		{
			Artist: "ANY_ARTIST",
			Price:  100.00,
			Title:  "ANY_TITLE",
		},
	}
	testCases := []struct {
		expectedResponse   entity.Response
		message            string
		mockMethodResponse []entity.Album
		mockMethodError    error
		name               string
	}{
		{
			expectedResponse: entity.Response{
				StatusCode: 200,
				Message:    "OK",
				Data:       expectedAlbums,
			},
			message:            "should return the OK Response",
			mockMethodResponse: expectedAlbums,
			mockMethodError:    nil,
			name:               "OK Response",
		},
		{
			expectedResponse: entity.Response{
				StatusCode: 500,
				Message:    "Internal Server Error",
				Data:       nil,
			},
			message:            "should return the Internal Server Error Response",
			mockMethodResponse: []entity.Album{},
			mockMethodError:    errors.New("Get Error"),
			name:               "Internal Server Error Response",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			albumRepository := new(albumRepositoryMock)
			albumHandler := handler.NewAlbumREST(albumRepository)

			albumRepository.On("GetAll").Return(testCase.mockMethodResponse, testCase.mockMethodError)

			actualResponse := albumHandler.GetAlbums()

			assert.Equal(t, actualResponse, testCase.expectedResponse, testCase.message)
		})
	}
}
