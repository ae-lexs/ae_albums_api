package handler

import (
	"testing"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/repository"
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
	expectedAlbum := entity.Album{
		Artist: "ANY_ARTIST",
		Price:  0.0,
		Title:  "ANY_TITLE",
	}
	albumRepository := &repository.PostgresAlbumRepository{
		Client: &FakeRepositoryClient{[]entity.Album{}},
	}

	if actualAlbum, _ := CreateAlbum(expectedAlbum, albumRepository); actualAlbum != expectedAlbum {
		t.Error("The actualAlbum is not equal to the expectedAlbum")
	}
}

func TestGetAlbums(t *testing.T) {
	expectedAlbum := entity.Album{
		Title:  "ANY_TITLE",
		Artist: "ANY_TITLE",
		Price:  10.0,
	}
	albumRepository := &repository.PostgresAlbumRepository{
		Client: &FakeRepositoryClient{[]entity.Album{}},
	}

	CreateAlbum(expectedAlbum, albumRepository)

	if actualAlbums, _ := GetAlbums(albumRepository); actualAlbums[0] != expectedAlbum {
		t.Error("The actualAlbums[0] is not equal to the expectedAlbum")
	}
}
