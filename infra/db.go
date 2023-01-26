package infra

import (
	"RegLog/app/response"
	"net/http"
)

type DB interface {
	Insert(user *response.Users, w http.ResponseWriter) error
	Get(user *response.Users, w http.ResponseWriter) (response.Users, error)
}
