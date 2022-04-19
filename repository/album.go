package repository

import "github.com/ae-lexs/ae_albums_api/entity"

type AlbumRepository interface {
	Create(entity.Album) entity.Album
	GetAll() []entity.Album
}

type RepositoryClient interface {
	Create(*entity.Album) entity.Album
	Find(*[]entity.Album) []entity.Album
}

type PostgresAlbumRepository struct {
	Client RepositoryClient
}

func (p *PostgresAlbumRepository) Create(album entity.Album) entity.Album {
	return p.Client.Create(&album)
}

func (p *PostgresAlbumRepository) GetAll() []entity.Album {
	var albums []entity.Album

	return p.Client.Find(&albums)
}