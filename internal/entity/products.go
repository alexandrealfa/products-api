package entity

import (
	"errors"
	"github.com/alexandrealfa/products-api/pkg/entity"
	"time"
)

var (
	errorIdIsEmpty      = errors.New("id is empty")
	errorNameIsEmpty    = errors.New("name is empty")
	errorPriceRequired  = errors.New("price is required")
	errorIdIsInvalid    = errors.New("id is invalid")
	errorNameIsRequired = errors.New("name is required")
	errorInvalidPrice   = errors.New("invalid price")
)

type Products struct {
	Id        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateProduct(name string, price float64) *Products {
	return &Products{
		entity.NewId(),
		name,
		price,
		time.Now(),
	}
}

func (p *Products) Validate() error {
	validateId := func(_ entity.ID, err error) bool { return err != nil }

	switch {
	default:
		return nil
	case p.Id.String() == "":
		return errorIdIsEmpty
	case p.Name == "":
		return errorNameIsEmpty
	case p.Price < 0:
		return errorInvalidPrice
	case p.Price == 0:
		return errorPriceRequired
	case len(p.Name) <= 1:
		return errorNameIsRequired
	case validateId(entity.ParseId(p.Id.String())):
		return errorIdIsInvalid
	}
}
