package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := CreateProduct("ProductNew", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.Id)
	assert.Equal(t, "ProductNew", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestWhenNameIsRequired(t *testing.T) {
	p, err := CreateProduct("", 10)

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, errorNameIsEmpty, err)
}

func TestWhenPriceIsRequired(t *testing.T) {
	p, err := CreateProduct("productValue", 0)

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, errorPriceRequired, err)
}

func TestWhenPriceIsInvalid(t *testing.T) {
	p, err := CreateProduct("product Invalid", -10)

	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, errorInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := CreateProduct("validProduct", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "validProduct")
	assert.Nil(t, p.Validate())
}
