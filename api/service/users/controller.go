package users

import "github.com/gin-gonic/gin"

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users",
	})
}
