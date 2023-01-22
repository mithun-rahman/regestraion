package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status   int     `json:"-"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	//Results  interface{} `json:"results"`
	//Count    int32       `json:"count" bson:"count"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ServeJSON serves json to http client
func (r *Response) ServeJSON(w http.ResponseWriter) error {
	resp := &Response{
		Next:     r.Next,
		Previous: r.Previous,
		Data:     r.Data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}
	return nil
}

// ServeJSON a utility func which serves json to http client
func ServeJSON(w http.ResponseWriter, status int, previous, next *string, message string, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := &Response{
		Status:   status,
		Next:     next,
		Previous: previous,
		//Results:  data,
		Data:    data,
		Message: message,
		//Count:   count,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}

	return nil
}
