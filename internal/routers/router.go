package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/pqt2p1/go-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
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
