package database

import (
	"errors"

	"github.com/joho/godotenv"
)

func EnvInit() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}
	return nil
}