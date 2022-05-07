package model

import (
	"github.com/ae-lexs/ae_albums_api/entity"
	"gorm.io/gorm"
)

type Album interface {
	Create(entity.Album) (entity.Album, error)
	Find() ([]entity.Album, error)
}

type albumPostgres struct {
	client *gorm.DB
}

func NewAlbumPostgres(client *gorm.DB) albumPostgres {
	return albumPostgres{
		client: client,
	}
}

func (model *albumPostgres) Create(album entity.Album) (entity.Album, error) {
	if err := model.client.Create(&album).Error; err != nil {
		return album, err
	}

	return album, nil
}

func (model *albumPostgres) Find() ([]entity.Album, error) {
	var albums []entity.Album

	if err := model.client.Find(&albums).Error; err != nil {
		return nil, err
	}

	result := model.client.Find(&albums)

	if err := result.Error; err != nil {
		return albums, err
	}

	return albums, nil
}
