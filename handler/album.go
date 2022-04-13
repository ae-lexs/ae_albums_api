package handler

import (
	"github.com/ae-lexs/ae_albums_api/entity"

	"github.com/ae-lexs/ae_albums_api/repository"
)

func GetAlbums(r repository.AlbumRepository) ([]entity.Album, error) {
	return r.GetAll(), nil
}

func CreateAlbum(a entity.Album, r repository.AlbumRepository) (entity.Album, error) {
	r.Create(a)

	return a, nil
}
