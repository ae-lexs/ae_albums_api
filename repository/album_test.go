package repository

import (
	"reflect"
	"testing"

	"github.com/ae-lexs/ae_albums_api/entity"
)

type FakeRepositoryClient struct {
	albums []entity.Album
}

func (f *FakeRepositoryClient) Create(album *entity.Album) entity.Album {
	f.albums = append(f.albums, *album)

	return *album
}

func (f *FakeRepositoryClient) Find(*[]entity.Album) []entity.Album {
	return f.albums
}

func TestCreateAlbum(t *testing.T) {
	albumRepository := &PostgresAlbumRepository{
		client: &FakeRepositoryClient{[]entity.Album{}},
	}
	expectedAlbum := entity.Album{
		Title:  "ANY_TITLE",
		Artist: "ANY_TITLE",
		Price:  10.0,
	}

	actualCreatedAlbum := albumRepository.Create(expectedAlbum)

	if actualCreatedAlbum != expectedAlbum {
		t.Error("The actual_created_album is not equal to the expectedAlbum")
	}
}

func TestGetAlbums(t *testing.T) {
	album := &PostgresAlbumRepository{
		client: &FakeRepositoryClient{[]entity.Album{}},
	}
	expectedAlbum := entity.Album{
		Title:  "ANY_TITLE",
		Artist: "ANY_TITLE",
		Price:  10.0,
	}
	expectedAlbums := []entity.Album{expectedAlbum}
	album.Create(expectedAlbum)

	actualAlbums := album.GetAll()

	if !reflect.DeepEqual(actualAlbums, expectedAlbums) {
		t.Error("The actualAlbums is not equal to the expectedAlbums")
	}
}
