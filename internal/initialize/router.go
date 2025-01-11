package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/pqt2p1/go-ecommerce-backend-api/internal/controller"
	"github.com/pqt2p1/go-ecommerce-backend-api/internal/middleware"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before AA")
		c.Next()
		fmt.Println("After AA")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before CC")
	c.Next()
	fmt.Println("After CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthenMiddleware())

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Ping)
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}

	v2 := r.Group("/v2/2024")
	{
		v2.GET("/ping", c.NewPongController().Ping)
		v2.GET("/user/1", c.NewUserController().GetUserByID)
	}
	return r
}
