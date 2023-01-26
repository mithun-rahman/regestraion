package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
		if err := json.NewEncoder(w).Encode(resp.Data.(Token)); err != nil {
			fmt.Println(v)
			return err
		}
	case Users:
		if err := json.NewEncoder(w).Encode(resp.Data.(Users)); err != nil {
			return err
		}
	}
	return nil
}
