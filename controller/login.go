package controller

import (
	"RegLog/model"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"net/http"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {

	var payload, fromDb model.Users

	err := json.NewDecoder(r.Body).Decode(&payload)
	resp := model.Response{500, "problem with internal serve", nil}

	if err != nil {
		resp.ServeJSON(w)
		return
	}

	fmt.Println(payload)

	logErr := c.DB.Where("user_name = ?", payload.User_name).First(&fromDb).Error

	if logErr != nil {
		resp := model.Response{404, "user not found", nil}
		resp.ServeJSON(w)
		return
	}

	fmt.Println(fromDb.User_name)

	ok, err := argon2id.ComparePasswordAndHash(payload.Password, fromDb.Password)

	if err != nil {
		resp.ServeJSON(w)
		return
	}
	if ok {
		token, err := c.GenerateToken(&model.Users{
			User_name: payload.User_name,
			Email:     payload.Email,
		})
		if err != nil {
			resp.ServeJSON(w)
			return
		}
		resp := model.Response{200, "", token}
		resp.ServeJSON(w)
	} else {
		resp.ServeJSON(w)
		return
	}
}
