package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	User_name string `json:"user_Name"`
	Password  string `json:",omitempty"`
	Email     string `json:"email"`
}

type User struct {
	User_name string `json:"user_name"`
	Email     string `json:"email"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	Refresh_token string `json:"refresh_token"`
}

type Error struct {
	Message string `json:"message"`
}

type Err struct {
	Message string `json:"message"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) ServeJSON(w http.ResponseWriter) error {

	resp := &Response{
		Data:    r.Data,
		Status:  r.Status,
		Message: r.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)

	if len(resp.Message) > 0 {
		msg := Err{resp.Message}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return err
		}
		return nil
	}

	switch v := resp.Data.(type) {
	case Token:
		fmt.Println(v)
		if err := json.NewEncoder(w).Encode(resp.Data.(Token)); err != nil {
			return err
		}
	case Users:
		user := User{resp.Data.(Users).User_name, resp.Data.(Users).Email}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			return err
		}
	}
	return nil
}
