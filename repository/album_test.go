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
