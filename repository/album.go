package repository

import (
	"github.com/ae-lexs/ae_albums_api/entity"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	Create(entity.Album) (entity.Album, error)
	GetAll() ([]entity.Album, error)
}

type PostgresAlbumRepository struct {
	Client *gorm.DB
}

func (repository *PostgresAlbumRepository) Create(album entity.Album) (entity.Album, error) {
	if err := repository.Client.Create(album).Error; err != nil {
		return album, err
	}

	return album, nil
}

func (repository *PostgresAlbumRepository) GetAll() ([]entity.Album, error) {
	var albums []entity.Album

	if err := repository.Client.Find(&albums).Error; err != nil {
		return nil, err
	}

	result := repository.Client.Find(&albums)

	if err := result.Error; err != nil {
		return albums, err
	}

	return albums, nil
}
