package app

import (
	"RegLog/middle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func brandsRouter(ctrl *BrandsController) http.Handler {
	r := chi.NewRouter()

	r.Use(middle.Limit)
	r.Use(middleware.DefaultLogger)

	r.Post("/reg", ctrl.Registration)
	r.Post("/log", ctrl.Login)
	r.Post("/refresh", ctrl.Refresh)
	return r
}
