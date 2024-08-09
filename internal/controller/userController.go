package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"masterbi1", "masterbi2"})
}

func (uc *UserController) Register(c *gin.Context) {
	email := c.PostForm("email")
	purpose := c.PostForm("purpose")

	code := uc.userService.Register(email, purpose)
	response.SuccessResponse(c, code, nil)
}
