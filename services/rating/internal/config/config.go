package config

import (
	"errors"
	"fmt"
	"os"
)

func GetConnectionString() (string, error) {
	// Берём обязательные части из переменных окружения
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return "", errors.New("DB_HOST not set")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return "", errors.New("DB_PORT not set")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return "", errors.New("DB_USER not set")
	}

	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return "", errors.New("DB_PASSWORD not set")
	}

	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return "", errors.New("DB_NAME not set")
	}

	// Формируем строку подключения
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbname,
	)

	return connStr, nil
}
