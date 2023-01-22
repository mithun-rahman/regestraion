package controller

import (
	"crypto/rsa"
	"gorm.io/gorm"
	"net/http"
)

type Controller struct {
	DB         *gorm.DB
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

type CallFunction interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	RefreshToken(http.ResponseWriter, *http.Request)
}
