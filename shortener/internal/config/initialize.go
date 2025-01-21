package config

import (
	"url-shortener/internal/controller"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
	"url-shortener/internal/store"

	"github.com/gorilla/mux"
)

func InitializeApp(router *mux.Router) {
	store := store.NewInMemoryStore()

	repo := repository.NewURLShortenerRepository(store)
	service := service.NewURLShortenerService(repo)
	controller := controller.NewURLShortenercontroller(service)

	controller.RegisterRoutes(router)
}
