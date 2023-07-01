package start

import (
	"github.com/gin-gonic/gin"
	route "github.com/lanngoen1996/golang-basic/start/routes"
)

func NewRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	route.HealthzRoute(v1)
	route.UserRouter(v1.Group("/users"))
}
