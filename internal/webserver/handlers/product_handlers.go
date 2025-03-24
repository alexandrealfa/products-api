package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alexandrealfa/products-api/internal/database"
	"github.com/alexandrealfa/products-api/internal/dto"
	"github.com/alexandrealfa/products-api/internal/entity"
	entityPKG "github.com/alexandrealfa/products-api/pkg/entity"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{db}
}

func (H *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductDTO
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newProduct, err := entity.CreateProduct(product.Name, product.Price)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = H.ProductDB.Create(newProduct); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	products, err := H.ProductDB.FindAll(1, 10, "asc")
	if err != nil {
		log.Println(err)
	}
	for _, products := range products {
		fmt.Println(products.Name, products.Id)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(newProduct.Id.String()))
}

func (H *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	product, err := H.ProductDB.FindById(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(product)
}

func (H *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var content entity.Products

	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	content.Id, err = entityPKG.ParseId(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w)
		return
	}
	if _, err := H.ProductDB.FindById(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	if err := H.ProductDB.Update(&content); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(content)
}

func (H *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := H.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := H.ProductDB.Delete(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (H *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	sort := chi.URLParam(r, "sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	products, err := H.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
