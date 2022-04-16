package repository

import (
	"github.com/micrease/gorme"
	"gorm.io/gorm"
	"meshop-api/app/model"
)

type Product struct {
	gorme.Repository[model.Product]
}

func NewProduct(db *gorm.DB) *Product {
	repo := new(Product)
	repo.SetDB(db)
	return repo
}
