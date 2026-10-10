package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/Adisey/car-rental-api/internal/handlers"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println(".env.local file not found")
	}
	ctx := context.Background()
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	if err := db.ConnectBun(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/cars", handlers.CarsMainHandler)
	http.HandleFunc("/cars/", handlers.CarByIDMainHandler)
	http.HandleFunc("/colors", handlers.ColorsMainHandler)
	http.HandleFunc("/colors/", handlers.ColorByIDMainHandler)
	http.HandleFunc("/brands", handlers.BrandsMainHandler)
	http.HandleFunc("/brands/", handlers.BrandByIDMainHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
