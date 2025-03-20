package database

import (
	"fmt"
	"github.com/alexandrealfa/products-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func generateConnection() (db *gorm.DB, err error) {
	dsn := "file::memory:"
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entity.Products{})
	return
}

func TestCreateNewProduct(t *testing.T) {
	db, err := generateConnection()

	if err != nil {
		t.Error(err)
	}

	product, er := entity.CreateProduct("default_product", 10.50)

	assert.NoError(t, er)

	productDB := NewProduct(db)

	if err = productDB.Create(product); err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, product.Id)

	var productValidate entity.Products

	if err = db.First(&productValidate, "id = ?", product.Id).Error; err != nil {
		t.Error(err)
	}

	assert.Equal(t, productValidate.Name, "default_product")
	assert.Equal(t, productValidate.Price, 10.50)
}

func TestAllProducts(t *testing.T) {
	db, err := generateConnection()
	assert.NoError(t, err)

	var productDb = NewProduct(db)

	for i := 1; i <= 20; i++ {
		price := rand.Float64() * 100
		product, er := entity.CreateProduct(fmt.Sprintf("product %d", i), price)
		assert.NoError(t, er)
		assert.NotNil(t, product.Name)
		assert.Equal(t, product.Price, price)
		db.Create(&product)
	}

	products, err := productDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "product 1")
	assert.Equal(t, products[9].Name, "product 10")

	products, err = productDb.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "product 11")
	assert.Equal(t, products[9].Name, "product 20")
}

func TestGetProductById(t *testing.T) {
	db, err := generateConnection()
	assert.NoError(t, err)

	productDb := NewProduct(db)

	product, err := entity.CreateProduct("new product", 1000.0)
	assert.NoError(t, err)

	if err = productDb.Create(product); err != nil {
		t.Error(err)
	}

	productByData, err := productDb.FindById(product.Id.String())
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.Name, productByData.Name)
	assert.Equal(t, product.Price, productByData.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := generateConnection()
	assert.NoError(t, err)

	product, err := entity.CreateProduct("MyProduct", 250.50)
	if err != nil {
		t.Error(err)
	}

	productDB := NewProduct(db)
	productDB.Create(product)

	NameUpdated := "New Name"
	product.Name = NameUpdated

	assert.NoError(t, productDB.Update(product))

	productByData, err := productDB.FindById(product.Id.String())

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, productByData.Name, NameUpdated)
}

func TestDeleteProduct(t *testing.T) {
	db, err := generateConnection()
	assert.NoError(t, err)

	product, err := entity.CreateProduct("product random", rand.Float64()*100)

	if err != nil {
		t.Error(err)
	}

	productDB := NewProduct(db)
	productDB.Create(product)

	assert.NoError(t, productDB.Delete(product.Id.String()))

	_, err = productDB.FindById(product.Id.String())
	assert.Error(t, err)
}
