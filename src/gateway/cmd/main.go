package main

import (
	"os"

	// "github.com/joho/godotenv"
	"github.com/spf13/viper"

	handler "github.com/Mamvriyskiy/lab2-template/src/gateway/handler"
	services "github.com/Mamvriyskiy/lab2-template/src/gateway/services"
	server "github.com/Mamvriyskiy/lab2-template/src/server"
)

func main() {
	services := services.NewService()
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		return
	}
}
