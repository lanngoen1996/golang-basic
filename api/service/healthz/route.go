package healthz

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	healthz := r.Group("healthz")
	healthz.Use()
	{
		healthz.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
