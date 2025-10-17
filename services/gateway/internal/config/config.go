package config

import (
	"errors"
	"gateway/models"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load(".env")
}

func GetConnectionString() (models.ServiceConf, error) {
	conf := models.ServiceConf{}
	key, flag := os.LookupEnv("LIB_SERVICE_URL")
	if !flag {
		return models.ServiceConf{}, errors.New("connection string not found")
	}
	key2, flag := os.LookupEnv("RATING_SERVICE_URL")
	if !flag {
		return models.ServiceConf{}, errors.New("connection string not found")
	}
	key3, flag := os.LookupEnv("RESERVATION_SERVICE_URL")
	if !flag {
		return models.ServiceConf{}, errors.New("connection string not found")
	}
	conf.LibraryURL = key
	conf.RatingURL = key2
	conf.ReservationURL = key3
	return conf, nil
}
