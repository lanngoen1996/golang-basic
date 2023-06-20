package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lanngoen1996/golang-basic/cors"
)

func main() {
	r := gin.Default()

	cors.NewApp(r)

	r.Run(cors.GetEnv().HOST + ":" + cors.GetEnv().PORT)
}
