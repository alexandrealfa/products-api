package database

import "github.com/alexandrealfa/products-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Products) error
	FindAll(page, limit int, sort string) ([]entity.Products, error)
	FindById(Id string) (*entity.Products, error)
	Update(product *entity.Products) error
	Delete(id string) error
}
