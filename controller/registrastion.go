package controller

import (
	"RegLog/model"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"net/http"
)

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {

	var payload, user model.Users

	err := json.NewDecoder(r.Body).Decode(&payload)
	resp := model.Response{500, "problem with internal serve", nil}

	if err != nil {
		resp.ServeJSON(w)
		return
	}

	fmt.Println(payload)

	regErr := c.DB.Where("user_name = ?", payload.User_name).First(&user).Error

	if regErr == nil {
		resp := model.Response{500, "user already exits", nil}
		resp.ServeJSON(w)
		return
	}

	userPassword := payload.Password
	hash, err := argon2id.CreateHash(userPassword, argon2id.DefaultParams)

	if err != nil {
		resp := model.Response{500, "problem with hash function", nil}
		resp.ServeJSON(w)
		return
	}

	payload.Password = hash

	er := c.DB.Create(&payload).Error

	fmt.Println(er)

	token, err := c.GenerateToken(&model.Users{
		User_name: payload.User_name,
		Email:     payload.Email,
	})

	resp = model.Response{200, "", token}
	resp.ServeJSON(w)
}
