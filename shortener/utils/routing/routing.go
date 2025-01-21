package routing

import (
	"url-shortener/utils/middleware"

	"github.com/gorilla/mux"
)

func NewDefaultRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.ReqLogger)
	return router
}
