package service

import (
	"math/rand"
	"strings"
	"time"

	"url-shortener/internal/command"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

var (
	HTTPS_PREFIX = "https://"
)

type URLShortenerService struct {
	repo *repository.URLShortenerRepository
}

func NewURLShortenerService(repo *repository.URLShortenerRepository) *URLShortenerService {
	return &URLShortenerService{repo: repo}
}

func (service *URLShortenerService) ShortenURL(cmd *command.CreateShortenURLCommand) (string, error) {
	alias := cmd.CustomAlias
	if cmd.CustomAlias == "" {
		alias = service.shortenURL()
	}

	data := model.URLData{
		ShortURL:         alias,
		LongURL:          cmd.LongURL,
		AccessCount:      0,
		AccessTimestamps: make([]string, 0),
	}

	err := service.repo.AddNewURL(&data)
	if err != nil {
		return "", err
	}

	return data.ShortURL, nil
}

func (service *URLShortenerService) DecodeShortURL(alias string) (string, error) {
	urlData, err := service.repo.GetURL(alias)
	if err != nil {
		return "", err
	}

	urlData.AccessCount++
	urlData.AccessTimestamps = append(urlData.AccessTimestamps, time.Now().String())

	longURL := urlData.LongURL
	longURL = strings.Trim(longURL, HTTPS_PREFIX)

	return longURL, nil
}

func (service *URLShortenerService) UpdateShortURL(cmd *command.UpdateShortURLCommand) error {
	return nil
}

func (service *URLShortenerService) GetAnalytics(alias string) (*model.URLData, error) {
	urlData, err := service.repo.GetURL(alias)
	if err != nil {
		return nil, err
	}

	return urlData, nil
}

func (service *URLShortenerService) shortenURL() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	var shortURL [6]byte

	// Seed the random number generator to ensure randomness
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 6; i++ {
		randomNumber := rand.Intn(len(charset))
		shortURL[i] = charset[randomNumber]
	}

	return string(shortURL[:])
}
