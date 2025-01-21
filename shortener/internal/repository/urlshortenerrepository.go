package repository

import (
	"url-shortener/internal/model"
	"url-shortener/internal/store"
)

type URLShortenerRepository struct {
	datastore store.Store
}

func NewURLShortenerRepository(datastore store.Store) *URLShortenerRepository {
	return &URLShortenerRepository{
		datastore: datastore,
	}
}

func (repo *URLShortenerRepository) AddNewURL(urlData *model.URLData) error {
	err := repo.datastore.Add(urlData)
	if err != nil {
		return err
	}

	return nil
}

func (repo *URLShortenerRepository) GetURL(alias string) (*model.URLData, error) {
	urlData, err := repo.datastore.Get(alias)
	if err != nil {
		return nil, err
	}

	return urlData, nil
}

func (repo *URLShortenerRepository) UpdateURL(data *model.URLData) error {
	err := repo.datastore.Update(data)
	if err != nil {
		return err
	}

	return nil
}

func (repo *URLShortenerRepository) DeleteURL(alias string) error {
	err := repo.datastore.Delete(alias)
	if err != nil {
		return err
	}

	return nil
}
