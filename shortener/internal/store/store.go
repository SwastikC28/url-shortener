package store

import (
	"strings"
	"url-shortener/internal/model"
	"url-shortener/utils/errors"
)

type Store interface {
	Get(string) (*model.URLData, error)
	Add(*model.URLData) error
	Update(*model.URLData) error
	Delete(string) error
	Exists(string) bool
}

type InMemoryStore struct {
	store map[string]*model.URLData
}

func NewInMemoryStore() Store {
	return &InMemoryStore{
		store: make(map[string]*model.URLData),
	}
}

func (inmemorystore *InMemoryStore) Add(data *model.URLData) error {
	exists := inmemorystore.Exists(data.ShortURL)
	if exists {
		return errors.NewValidationError("resource already exists in the system")
	}

	inmemorystore.store[strings.ToLower(data.ShortURL)] = data
	return nil
}

func (inmemorystore *InMemoryStore) Delete(alias string) error {
	exists := inmemorystore.Exists(alias)
	if !exists {
		return errors.NewResourceNotFoundError("resource does not exists in the system")
	}

	delete(inmemorystore.store, alias)
	return nil
}

func (inmemorystore *InMemoryStore) Update(data *model.URLData) error {
	exists := inmemorystore.Exists(data.ShortURL)
	if !exists {
		return errors.NewResourceNotFoundError("resource does not exists in the system")
	}

	inmemorystore.store[data.ShortURL] = data
	return nil
}

func (inmemorystore *InMemoryStore) Exists(alias string) bool {
	_, exists := inmemorystore.store[alias]
	return exists
}

func (inmemorystore *InMemoryStore) Get(alias string) (*model.URLData, error) {
	data, exists := inmemorystore.store[alias]
	if !exists {
		return nil, errors.NewResourceNotFoundError("resource does not exists in the system")
	}

	return data, nil
}
