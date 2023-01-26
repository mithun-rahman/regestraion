package app

import (
	"RegLog/app/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter(brndsCtrl *BrandsController) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.NotFound(NotFoundHandler)
	r.MethodNotAllowed(MethodNotAllowed)

	r.Mount("/user", brandsRouter(brndsCtrl))

	return r
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{406, "method not acceptable", nil}
	resp.ServeJSON(w)
}

// MethodNotAllowed handles when no routes match
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{405, "method not allowed", nil}
	resp.ServeJSON(w)
}
