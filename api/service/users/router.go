package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(r *gin.RouterGroup, DB *gorm.DB) {
	users := r.Group("users")

	controller := NewUserController(DB)

	users.Use()
	{
		users.POST("/create", controller.CreateUser)
		users.GET("/:id", controller.FindUser)
		users.PUT("/:id", controller.UpdateUser)
		users.DELETE("/:id", controller.Delete)
		users.GET("/", controller.GetUsers)
	}
}
