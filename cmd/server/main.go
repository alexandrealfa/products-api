package main

import (
	"fmt"
	"github.com/alexandrealfa/products-api/configs"
	"log"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.DBName)
}
