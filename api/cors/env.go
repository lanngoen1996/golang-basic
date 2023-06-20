package cors

import "os"

type Env struct {
	HOST string `json:"host"`
	PORT string `json:"port"`
}

func GetEnv() Env {
	return Env{
		HOST: os.Getenv("HOST"),
		PORT: os.Getenv("PORT"),
	}
}
