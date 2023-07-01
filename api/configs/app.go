package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/lanngoen1996/golang-basic/core"
	"github.com/lanngoen1996/golang-basic/database"
	"github.com/lanngoen1996/golang-basic/start"
)

func NewApp(r *gin.Engine) {
	core.ConnectDB()
	database.Migration()
	start.NewRouter(r)
}
