package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/pqt2p1/go-ecommerce-backend-api/global"
	"github.com/pqt2p1/go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middleware
	r.Use() // logging
	r.Use() // cross
	r.Use() // limiter global

	manageRouter := routers.RouterGroupApp.ManageRouterGroup
	userRouter := routers.RouterGroupApp.UserRouterGroup

	MainGroup := r.Group("v1/2024")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitProductRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}

	return r
}
