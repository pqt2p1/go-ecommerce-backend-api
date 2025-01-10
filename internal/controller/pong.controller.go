package controller

import "github.com/gin-gonic/gin"

type PongController struct {}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
