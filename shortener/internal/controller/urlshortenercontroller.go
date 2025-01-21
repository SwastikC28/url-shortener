package controller

import (
	"fmt"
	"net/http"
	"url-shortener/internal/command"
	"url-shortener/internal/dto"
	"url-shortener/internal/service"
	"url-shortener/utils/routing"
	"url-shortener/utils/web"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type URLShortenerController struct {
	service *service.URLShortenerService
}

func NewURLShortenercontroller(service *service.URLShortenerService) routing.Controller {
	return &URLShortenerController{
		service: service,
	}
}

func (controller *URLShortenerController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/shorten", controller.shortenURL).Methods(http.MethodPost)
	router.HandleFunc("/{alias}", controller.redirect).Methods(http.MethodGet)
	router.HandleFunc("/update/{alias}", controller.updateURL).Methods(http.MethodPut)
	router.HandleFunc("/analytics/{alias}", controller.analytics).Methods(http.MethodGet)
}

func (controller *URLShortenerController) shortenURL(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context()).With().Str("action", "shortenURL").Logger()

	cmd := command.CreateShortenURLCommand{}

	err := web.UnmarshalJSON(r, &cmd)
	if err != nil {
		logger.Err(err).Msg("Error while unmarshalling JSON")
		web.RespondJSON(w, 400, err)
		return
	}

	shortURL, err := controller.service.ShortenURL(&cmd)
	if err != nil {
		logger.Err(err).Msg("Error while shortening URL")
		web.RespondJSON(w, 400, err)
		return
	}

	shortenURLDTO := dto.ShortenURLDTO{
		ShortURL: shortURL,
	}

	web.RespondJSON(w, http.StatusCreated, shortenURLDTO)
}

func (controller *URLShortenerController) redirect(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context()).With().Str("action", "redirect").Logger()

	vars := mux.Vars(r)
	shortURL := vars["alias"]

	url, err := controller.service.DecodeShortURL(shortURL)
	if err != nil {
		logger.Err(err).Msg("Unable to decode URL")
		web.RespondJSON(w, 400, err)
		return
	}

	longURL := fmt.Sprintf("https://%s", url)
	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

func (controller *URLShortenerController) updateURL(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context()).With().Str("action", "redirect").Logger()

	cmd := &command.UpdateShortURLCommand{}

	err := web.UnmarshalJSON(r, cmd)
	if err != nil {
		logger.Err(err).Msg("Error while unmarshalling JSON")
		web.RespondJSON(w, 400, err)
		return
	}

	err = controller.service.UpdateShortURL(cmd)
	if err != nil {
		logger.Err(err).Msg("Error while updating url")
		web.RespondJSON(w, 400, err)
		return
	}

	web.RespondJSON(w, 200, nil)
}

func (controller *URLShortenerController) analytics(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context()).With().Str("action", "analytics").Logger()

	vars := mux.Vars(r)
	alias := vars["alias"]

	urlData, err := controller.service.GetAnalytics(alias)
	if err != nil {
		logger.Err(err).Msg("Error while getting data")
		web.RespondJSON(w, 400, err.Error())
		return
	}

	web.RespondJSON(w, 200, urlData)
}
