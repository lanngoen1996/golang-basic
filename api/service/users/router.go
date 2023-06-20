package users

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	users := r.Group("users")
	users.Use()
	{
		users.GET("/", getUsers)
	}
}
