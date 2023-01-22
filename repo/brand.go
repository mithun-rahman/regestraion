package repo

import (
	"RegLog/infra"
)

type BrandRepo interface {
	Repo
}

type PostresBrand struct {
	db infra.DB
}

// NewBrand returns new brand repo
func NewBrand(table string, db infra.DB) BrandRepo {
	return &PostresBrand{
		db: db,
	}
}
