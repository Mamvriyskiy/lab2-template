package main

import (
	"os"

	"log"

	handler "github.com/Mamvriyskiy/lab2-template/src/flight/handler"
	repo "github.com/Mamvriyskiy/lab2-template/src/flight/repository"
	services "github.com/Mamvriyskiy/lab2-template/src/flight/services"
	server "github.com/Mamvriyskiy/lab2-template/src/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found: %v", err)
	}

	db, err := repo.NewPostgresDB(&repo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		log.Fatal("Error connect db:", err.Error())
		return
	}

	repos := repo.NewRepository(db)
	service := services.NewServices(repos)
	handlers := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run("8060", handlers.InitRouters()); err != nil {
		log.Fatal("Error occurred while running http server:", err.Error())
		return
	}
}
