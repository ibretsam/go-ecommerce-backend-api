package manager

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// Public Route
	ur := repo.NewUserRepo()
	us := service.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)

	publicUserRouter := Router.Group("/user")
	{
		publicUserRouter.POST("/register", userHandlerNonDependency.Register)
		publicUserRouter.POST("/otp")
	}

	// Private Route
	privateUserRouter := Router.Group("/admin/user")
	{
		privateUserRouter.GET("/activateUser")
	}
}
