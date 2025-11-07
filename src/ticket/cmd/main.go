package main

import (
	"log"
	"os"

	server "github.com/Mamvriyskiy/lab2-template/src/server"
	handler "github.com/Mamvriyskiy/lab2-template/src/ticket/handler"
	repo "github.com/Mamvriyskiy/lab2-template/src/ticket/repository"
	service "github.com/Mamvriyskiy/lab2-template/src/ticket/services"
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
		return
	}

	repos := repo.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8070", handlers.InitRouters()); err != nil {
		return
	}
}
