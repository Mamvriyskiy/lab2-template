package main
package main

import (
	"os"

	// "github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	handler "github.com/Mamvriyskiy/lab2-template/src/gateway/handler"
	services "github.com/Mamvriyskiy/lab2-template/src/gateway/services"
	server "github.com/Mamvriyskiy/lab2-template/src/server"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		return
	}

	db, err := repo.NewPostgresDB(&repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		return
	}

	repos := repo.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		return
	}
}
