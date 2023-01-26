package repo

import (
	"RegLog/app/response"
	"RegLog/infra"
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type BrandRepo interface {
	Repo
	CreateUser(user response.Users, w http.ResponseWriter) error
	GetUser(user response.Users, w http.ResponseWriter) error
	RefreshToken(payload response.RefreshToken, w http.ResponseWriter) error
}

type PostresBrand struct {
	db         infra.DB
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// NewBrand returns new brand repo
func NewBrand(db infra.DB, PrivateKey *rsa.PrivateKey, PublicKey *rsa.PublicKey) BrandRepo {
	return &PostresBrand{
		db:         db,
		PrivateKey: PrivateKey,
		PublicKey:  PublicKey,
	}
}

func (p *PostresBrand) CreateUser(user response.Users, w http.ResponseWriter) error {

	regErr := p.db.Insert(&user, w)
	if regErr != nil {
		return regErr
	}
	token, err := p.GenerateToken(&response.Users{
		User_name: user.User_name,
		Email:     user.Email,
	})
	if err != nil {
		resp := response.Response{500, "internal server problem", nil}
		resp.ServeJSON(w)
		return err
	}
	resp := response.Response{200, "", token}
	resp.ServeJSON(w)
	return nil
	return nil
}

func (p *PostresBrand) GetUser(user response.Users, w http.ResponseWriter) error {
	data, logErr := p.db.Get(&user, w)
	if logErr != nil {
		return logErr
	}
	token, err := p.GenerateToken(&response.Users{
		User_name: data.User_name,
		Email:     data.Email,
	})
	if err != nil {
		resp := response.Response{500, "internal server problem", nil}
		resp.ServeJSON(w)
		return err
	}
	resp := response.Response{200, "", token}
	resp.ServeJSON(w)
	return nil
}

func (p *PostresBrand) RefreshToken(payload response.RefreshToken, w http.ResponseWriter) error {

	resp := response.Response{500, "internal problem", nil}
	tokenString := payload.Refresh_token
	claims := jwt.MapClaims{}
	refresh_token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return p.PublicKey, nil
	})
	if err != nil {
		resp.ServeJSON(w)
		return nil
	}

	if !refresh_token.Valid {
		resp := response.Response{500, "invalid token", nil}
		resp.ServeJSON(w)
		return nil
	}

	user := response.Users{}
	for key, val := range claims {
		if strVal, ok := val.(string); ok && key == "email" {
			user.Email = strVal
		} else if strVal, ok := val.(string); ok && key == "user_name" {
			user.User_name = strVal
		}
	}

	new_token, err := p.GenerateToken(&response.Users{
		User_name: user.User_name,
		Email:     user.Email,
	})

	if err != nil {
		resp.ServeJSON(w)
		return nil
	}
	resp = response.Response{200, "", new_token}
	resp.ServeJSON(w)
	return nil
}
