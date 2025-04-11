package main

import (
	"fmt"
	"github.com/alexandrealfa/products-api/configs"
	"github.com/alexandrealfa/products-api/internal/database"
	"github.com/alexandrealfa/products-api/internal/entity"
	"github.com/alexandrealfa/products-api/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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
	userHandler := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExpiresIn)

	c := chi.NewRouter()
	c.Use(middleware.Logger)

	c.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		c.Post("/", productHandler.CreateProduct)
		c.Get("/{id}", productHandler.GetProduct)
		c.Get("/{page}/{limit}/{sort}", productHandler.GetProducts)
		c.Put("/{id}", productHandler.UpdateProduct)
		c.Delete("/{id}", productHandler.DeleteProduct)
	})

	c.Route("/users", func(r chi.Router) {
		c.Post("/", userHandler.CreateUser)
		c.Post("/generate-token", userHandler.GetJWT)
	})

	log.Println(http.ListenAndServe(":8002", c))
}
