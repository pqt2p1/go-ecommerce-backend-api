package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Ping)
		v1.POST("/ping", Ping)
		v1.PUT("/ping", Ping)
	}

	v2 := r.Group("/v2/2024")
	{
		v2.GET("/ping", Ping)
		v2.POST("/ping", Ping)
		v2.PUT("/ping", Ping)
	}
    return r
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
