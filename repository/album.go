package repository

import (
	"errors"
	"log"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/model"
)

var CreateError = errors.New("Album Repository Create Error")

type Album interface {
	Create(artist string, price float64, title string) (entity.Album, error)
	// GetAll() ([]entity.Album, error)
}

type album struct {
	model model.Album
}

func NewAlbum(model model.Album) album {
	return album{
		model: model,
	}
}

func (repository *album) Create(artist string, price float64, title string) (entity.Album, error) {
	createdAlbum, err := repository.model.Create(entity.Album{
		Artist: artist,
		Price:  price,
		Title:  title,
	})

	if err != nil {
		log.Print(err)
	}

	return createdAlbum, CreateError
}
