package main

import (
	"fmt"
	"github.com/alexandrealfa/products-api/configs"
	"github.com/alexandrealfa/products-api/internal/database"
	"github.com/alexandrealfa/products-api/internal/entity"
	"github.com/alexandrealfa/products-api/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.DBName)

	db, err := gorm.Open(sqlite.Open("productsdb"), &gorm.Config{})
	if err := db.AutoMigrate(&entity.Products{}, &entity.User{}); err != nil {
		log.Fatal(err)
	}
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Post("/products", productHandler.CreateProduct)
	c.Get("/product/{id}", productHandler.GetProduct)
	c.Get("/products/{page}/{limit}/{sort}", productHandler.GetProducts)
	c.Put("/products/{id}", productHandler.UpdateProduct)
	c.Delete("/products/{id}", productHandler.DeleteProduct)

	c.Post("/users", userHandler.CreateUser)
	log.Println(http.ListenAndServe(":8002", c))
}
