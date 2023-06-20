package cors

import "os"

type Env struct {
	HOST        string
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

func GetEnv() *Env {
	return &Env{
		HOST:        os.Getenv("HOST"),
		PORT:        os.Getenv("PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
}
