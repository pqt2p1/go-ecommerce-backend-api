package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pqt2p1/go-ecommerce-backend-api/internal/service"
	"github.com/pqt2p1/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, response.ErrCodeSuccess, []string{"user1", "user2"})
}
