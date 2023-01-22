package controller

import (
	"RegLog/model"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	str := r.Header.Get("Authorization")
	tokenString := strings.TrimSpace(strings.TrimPrefix(str, "Bearer"))

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return c.PublicKey, nil
	})
	resp := model.Response{500, "invalid token", nil}
	if err != nil {
		fmt.Println("problem is 21")
		resp.ServeJSON(w)
		return
	}

	if !token.Valid {
		fmt.Println("problem 27")
		resp.ServeJSON(w)
		return
	}

	// do something with decoded claims
	payload := model.Users{}
	for key, val := range claims {
		if strVal, ok := val.(string); ok && key == "email" {
			payload.Email = strVal
		} else if strVal, ok := val.(string); ok && key == "user_name" {
			payload.User_name = strVal
		}
	}

	var fromDB model.Users
	errLog := c.DB.Where("user_name = ?", payload.User_name).First(&fromDB).Error
	if errLog != nil {
		resp := model.Response{404, "user not found", nil}
		resp.ServeJSON(w)
		return
	}
	resp = model.Response{200, "", fromDB}
	resp.ServeJSON(w)
}
