package app

import (
	"RegLog/app/response"
	"RegLog/service"
	"encoding/json"
	"net/http"
)

type BrandsController struct {
	svc service.BrandService
}

func NewBrandsController(svc service.BrandService) *BrandsController {
	return &BrandsController{
		svc: svc,
	}
}

func (b *BrandsController) Registration(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{500, "", "server error"}
	var payload response.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		resp.ServeJSON(w)
		return
	}
	b.svc.CreateUser(payload, w)
}

func (b *BrandsController) Login(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{500, "", "server error"}
	var payload response.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		resp.ServeJSON(w)
		return
	}
	b.svc.GetUser(payload, w)
}

func (b *BrandsController) Refresh(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{500, "", "server error"}
	var payload response.RefreshToken
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		resp.ServeJSON(w)
		return
	}
	b.svc.RefreshToken(payload, w)
}
