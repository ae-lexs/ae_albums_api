package repository

import (
	"errors"
	"log"

	"github.com/ae-lexs/ae_albums_api/entity"
	"github.com/ae-lexs/ae_albums_api/model"
)

var (
	CreateError  = errors.New("Album Repository Create Error")
	GetAllError  = errors.New("Album Repository GetAll Error")
	GetByIDError = errors.New("Album Repository GetByIDError Error")
)

type Album interface {
	Create(artist string, price float64, title string) (entity.Album, error)
	GetAll() ([]entity.Album, error)
	GetByID(id string) ([]entity.Album, error)
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
		log.Printf("Album Repository Create, DB Client Error: %v", err)

		return entity.Album{}, CreateError
	}

	return createdAlbum, nil
}

func (respository *album) GetAll() ([]entity.Album, error) {
	foundAlbums, err := respository.model.Find()

	if err != nil {
		log.Printf("Album Repository Get, DB Client Error: %v", err)

		return foundAlbums, GetAllError
	}

	return foundAlbums, nil
}

func (respository *album) GetByID(id string) (entity.Album, error) {
	foundAlbum, err := respository.model.FindByID(id)

	if err != nil {
		log.Printf("Album Repository GetByID, DB Client Error: %v", err)

		return foundAlbum, GetByIDError
	}

	return foundAlbum, nil
}
