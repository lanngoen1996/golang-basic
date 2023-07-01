package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lanngoen1996/golang-basic/configs"
	"github.com/lanngoen1996/golang-basic/core"
)

func main() {
	r := gin.Default()
	configs.NewApp(r)
	r.Run(core.GetEnv().HOST + ":" + core.GetEnv().PORT)
}
