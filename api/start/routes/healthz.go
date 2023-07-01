package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/lanngoen1996/golang-basic/controllers"
)

func HealthzRoute(r *gin.RouterGroup) {
	c := controller.NewHealthzController()

	r.GET("/healthz", c.Healthz)
}
