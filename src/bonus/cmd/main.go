package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"

	handler "github.com/Mamvriyskiy/lab2-template/src/bonus/handler"
	services "github.com/Mamvriyskiy/lab2-template/src/bonus/services"
	server "github.com/Mamvriyskiy/lab2-template/src/server"
	repo "github.com/Mamvriyskiy/lab2-template/src/bonus/repository"
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
	if err := srv.Run("8050", handlers.InitRouters()); err != nil {
		log.Fatal("Failed to start server: ", err)
		return
	}
}
