package controller

import (
	"RegLog/model"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"
)

func (c *Controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	payload := model.RefreshToken{}

	resp := model.Response{500, "invalid token", nil}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		resp.ServeJSON(w)
		return
	}

	tokenString := payload.Refresh_token

	claims := jwt.MapClaims{}
	refresh_token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		verifyBytes, _ := ioutil.ReadFile("./app.rsa.pub")
		verifyKey, _ := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

		return verifyKey, nil
	})
	if err != nil {
		resp.ServeJSON(w)
		return
	}

	if !refresh_token.Valid {
		resp.ServeJSON(w)
		return
	}

	user := model.Users{}
	for key, val := range claims {
		if strVal, ok := val.(string); ok && key == "email" {
			user.Email = strVal
		} else if strVal, ok := val.(string); ok && key == "user_name" {
			user.User_name = strVal
		}
	}

	new_token, err := c.GenerateToken(&model.Users{
		User_name: user.User_name,
		Email:     user.Email,
	})

	if err != nil {
		resp.ServeJSON(w)
		return
	}

	resp = model.Response{200, "", new_token}
	resp.ServeJSON(w)
}
