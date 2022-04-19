package repository

import "github.com/ae-lexs/ae_albums_api/entity"

type FakeRepositoryClient struct {
	Albums []entity.Album
}

func (client *FakeRepositoryClient) Create(album entity.Album) error {
	client.Albums = append(client.Albums, album)

	return nil
}

func (client *FakeRepositoryClient) Find([]entity.Album) ([]entity.Album, error) {
	return client.Albums, nil
}

type FakeAlbumRepository struct {
	Client FakeRepositoryClient
}

func (repository *FakeAlbumRepository) Create(album entity.Album) (entity.Album, error) {
	if err := repository.Client.Create(album); err != nil {
		return album, err
	}

	return album, nil
}

func (repository *FakeAlbumRepository) GetAll() ([]entity.Album, error) {
	var albums []entity.Album

	foundAlbums, err := repository.Client.Find(albums)

	return foundAlbums, err
}
