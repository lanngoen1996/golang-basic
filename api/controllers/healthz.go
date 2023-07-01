package controller

import "github.com/gin-gonic/gin"

type HealthzControllerInterface interface {
	Healthz(ctx *gin.Context)
}

type HealthzController struct{}

func NewHealthzController() HealthzControllerInterface {
	return &HealthzController{}
}

func (c *HealthzController) Healthz(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
