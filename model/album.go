package model

import (
	"gorm.io/gorm"
)

type AlbumEntity struct {
	gorm.Model
	Artist string  `json:"artist"`
	Price  float64 `json:"price,string"`
	Title  string  `json:"title"`
}

type Album interface {
	Create(AlbumEntity) (AlbumEntity, error)
	Find() ([]AlbumEntity, error)
}

type albumPostgres struct {
	client *gorm.DB
}

func NewAlbumPostgres(client *gorm.DB) albumPostgres {
	return albumPostgres{
		client: client,
	}
}

func (model *albumPostgres) Create(album AlbumEntity) (AlbumEntity, error) {
	if err := model.client.Create(&album).Error; err != nil {
		return album, err
	}

	return album, nil
}

func (model *albumPostgres) Find() ([]AlbumEntity, error) {
	var albums []AlbumEntity

	if err := model.client.Find(&albums).Error; err != nil {
		return nil, err
	}

	result := model.client.Find(&albums)

	if err := result.Error; err != nil {
		return albums, err
	}

	return albums, nil
}
