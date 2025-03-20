package database

import (
	"github.com/alexandrealfa/products-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{db}
}

func (P *Product) Create(product *entity.Products) error {
	return P.DB.Create(product).Error
}

func (P *Product) FindById(Id string) (product *entity.Products, err error) {

	err = P.DB.First(&product, "id = ?", Id).Error

	return
}

func (P *Product) Update(product *entity.Products) error {
	_, err := P.FindById(product.Id.String())
	if err != nil {
		return err
	}

	return P.DB.Save(product).Error
}

func (P *Product) Delete(Id string) error {
	product, err := P.FindById(Id)

	if err != nil {
		return err
	}

	return P.DB.Delete(product).Error
}

func (P *Product) FindAll(page, limit int, sort string) (products []entity.Products, err error) {
	var myProducts []entity.Products

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = P.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&myProducts).Error
	} else {
		err = P.DB.Order("created_at " + sort).Find(&myProducts).Error
	}

	return myProducts, err
}

func GetCount(str string) (count int) {
	vowels := map[string]int{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0}
	for i := range str {
		if _, ok := vowels[string(str[i])]; ok {
			count++
		}
	}

	return
}
