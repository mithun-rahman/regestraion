package service

import (
	"RegLog/app/response"
	"RegLog/repo"
	"net/http"
)

type BrandService interface {
	CreateUser(user response.Users, w http.ResponseWriter) error
	GetUser(user response.Users, w http.ResponseWriter) error
	RefreshToken(payload response.RefreshToken, w http.ResponseWriter) error
}

type Brand struct {
	brandRepo repo.BrandRepo
}

func NewBrand(brandRepo repo.BrandRepo) BrandService {
	return &Brand{
		brandRepo: brandRepo,
	}
}

func (b *Brand) CreateUser(user response.Users, w http.ResponseWriter) error {
	if IsValid(user) == false {
		resp := response.Response{300, "input error", nil}
		resp.ServeJSON(w)
		return nil
	}
	b.brandRepo.CreateUser(user, w)
	return nil
}

func (b *Brand) GetUser(user response.Users, w http.ResponseWriter) error {
	b.brandRepo.GetUser(user, w)
	return nil
}

func (b *Brand) RefreshToken(payload response.RefreshToken, w http.ResponseWriter) error {
	b.brandRepo.RefreshToken(payload, w)
	return nil
}
