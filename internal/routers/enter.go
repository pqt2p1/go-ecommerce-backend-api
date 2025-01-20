package routers

import (
	"github.com/pqt2p1/go-ecommerce-backend-api/internal/routers/manage"
	"github.com/pqt2p1/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	UserRouterGroup   user.UserRouterGroup
	ManageRouterGroup manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
