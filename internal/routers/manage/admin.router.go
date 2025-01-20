package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	AdminRouterPublic := Router.Group("/admin") 
	{
		AdminRouterPublic.POST("/login")
	}

	// private router
	AdminRouterPrivate := Router.Group("/admin/user")
	{
		AdminRouterPrivate.POST("/active")
	}
	// AdminRouterPrivate.Use(Limiter())
	// AdminRouterPrivate.Use(Authen())
	// AdminRouterPrivate.Use(Permission())

}
