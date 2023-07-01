package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/lanngoen1996/golang-basic/controllers"
)

func UserRouter(r *gin.RouterGroup) {
	c := controller.NewUserController()

	r.POST("/create", c.CreateUser)
	r.GET("/:id", c.FindUser)
	r.PUT("/:id", c.UpdateUser)
	r.DELETE("/:id", c.Delete)
	r.GET("/", c.GetUsers)
}
