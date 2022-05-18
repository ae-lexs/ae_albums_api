package model

import (
	"github.com/ae-lexs/ae_albums_api/entity"
	"gorm.io/gorm"
)

type Album interface {
	Create(album entity.Album) (entity.Album, error)
	Find() ([]entity.Album, error)
	FindByID(id string) (entity.Album, error)
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
	var foundAlbums []entity.Album

	if err := model.client.Find(&foundAlbums).Error; err != nil {
		return foundAlbums, err
	}

	return foundAlbums, nil
}

func (model *albumPostgres) FindByID(id string) (entity.Album, error) {
	var foundAlbum entity.Album

	if err := model.client.First(&foundAlbum, id).Error; err != nil {
		return foundAlbum, err
	}

	return foundAlbum, nil
}
