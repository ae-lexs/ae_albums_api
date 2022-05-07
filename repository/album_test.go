package repository_test

import (
	"testing"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ModelMock struct {
	mock.Mock
}

func (mock *ModelMock) Create(album entity.Album) (entity.Album, error) {
	args := mock.Called(album)

	return album, args.Error(1)
}

func (mock *ModelMock) Find() ([]entity.Album, error) {
	args := mock.Called()

	return args.Get(0).([]entity.Album), args.Error(1)
}

func TestCreate(t *testing.T) {
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  100.0,
		Title:  "ANY_TITLE",
	}
	testCases := []struct {
		expectedAlbum entity.Album
		artist        string
		expectedError error
		price         float64
		title         string
		mockInput     entity.Album
		name          string
	}{
		{
			expectedAlbum: expectedAlbum,
			artist:        expectedAlbum.Artist,
			expectedError: nil,
			price:         expectedAlbum.Price,
			title:         expectedAlbum.Title,
			mockInput:     expectedAlbum,
			name:          "Create Album Without Error",
		},
		{
			expectedAlbum: entity.Album{},
			artist:        expectedAlbum.Artist,
			expectedError: repository.CreateError,
			price:         expectedAlbum.Price,
			title:         expectedAlbum.Title,
			mockInput:     expectedAlbum,
			name:          "Create Album With Error",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			modelMock := new(ModelMock)
			respository := repository.NewAlbum(modelMock)
			modelMock.On("Create", testCase.mockInput).Return(testCase.expectedAlbum, testCase.expectedError)

			actualAlbum, actualError := respository.Create(
				testCase.artist,
				testCase.price,
				testCase.title,
			)

			assert.True(t, modelMock.AssertExpectations(t))
			assert.Equal(t, actualAlbum, testCase.expectedAlbum)
			assert.Equal(t, actualError, testCase.expectedError)
		})
	}
}

func TestGetAll(t *testing.T) {
	expectedFoundAlbums := []entity.Album{
		{
			Artist: "ANY_ARTIST",
			Price:  100.0,
			Title:  "ANY_TITLE",
		},
	}
	testCases := []struct {
		expectedError  error
		expectedAlbums []entity.Album
		message        string
		name           string
	}{
		{
			expectedError:  nil,
			expectedAlbums: expectedFoundAlbums,
			name:           "Get All Albums Without Error",
		},
		{
			expectedError:  repository.GetAllError,
			expectedAlbums: []entity.Album{},
			name:           "Get All Albums With Error",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			modelMock := new(ModelMock)
			respository := repository.NewAlbum(modelMock)

			modelMock.On("Find").Return(testCase.expectedAlbums, testCase.expectedError)

			actualAlbums, actualError := respository.GetAll()

			assert.True(t, modelMock.AssertExpectations(t))
			assert.ElementsMatch(t, actualAlbums, testCase.expectedAlbums)
			assert.Equal(t, actualError, testCase.expectedError)
		})
	}
}
