package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Ping(c *gin.Context) {
	fmt.Println("-------> My handler")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
