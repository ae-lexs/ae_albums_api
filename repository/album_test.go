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
	modelMock := new(ModelMock)
	respository := repository.NewAlbum(modelMock)
	modelMock.On("Create", expectedAlbum).Return(expectedAlbum, nil)

	respository.Create(expectedAlbum.Artist, expectedAlbum.Price, expectedAlbum.Title)

	assert.True(t, modelMock.AssertExpectations(t))
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
		expectedError      error
		expectedResponse   []entity.Album
		message            string
		mockMethodResponse []entity.Album
		name               string
	}{
		{
			expectedError:      nil,
			expectedResponse:   expectedFoundAlbums,
			mockMethodResponse: expectedFoundAlbums,
			name:               "Get All Albums Without Error",
		},
		{
			expectedError:      repository.GetAllError,
			expectedResponse:   []entity.Album{},
			mockMethodResponse: []entity.Album{},
			name:               "Get All Albums With Error",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			modelMock := new(ModelMock)
			respository := repository.NewAlbum(modelMock)

			modelMock.On("Find").Return(testCase.mockMethodResponse, testCase.expectedError)

			actualAlbums, actualError := respository.GetAll()

			assert.True(t, modelMock.AssertExpectations(t))
			assert.ElementsMatch(t, actualAlbums, testCase.expectedResponse)
			assert.Equal(t, actualError, testCase.expectedError)
		})
	}
}
