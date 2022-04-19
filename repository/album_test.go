package repository

import (
	"reflect"
	"testing"

	"github.com/ae-lexs/ae_albums_api/entity"
)

func TestCreateAlbum(t *testing.T) {
	albumRepository := FakeAlbumRepository{
		Client: FakeRepositoryClient{[]entity.Album{}},
	}
	expectedAlbum := entity.Album{
		Title:  "ANY_TITLE",
		Artist: "ANY_TITLE",
		Price:  10.0,
	}

	actualCreatedAlbum, _ := albumRepository.Create(expectedAlbum)

	if actualCreatedAlbum != expectedAlbum {
		t.Error("The actual_created_album is not equal to the expectedAlbum")
	}
}

func TestGetAlbums(t *testing.T) {
	albumRepository := FakeAlbumRepository{
		Client: FakeRepositoryClient{[]entity.Album{}},
	}
	expectedAlbum := entity.Album{
		Title:  "ANY_TITLE",
		Artist: "ANY_TITLE",
		Price:  10.0,
	}
	expectedAlbums := []entity.Album{expectedAlbum}
	albumRepository.Create(expectedAlbum)

	actualAlbums, _ := albumRepository.GetAll()

	if !reflect.DeepEqual(actualAlbums, expectedAlbums) {
		t.Error("The actualAlbums is not equal to the expectedAlbums")
	}
}
