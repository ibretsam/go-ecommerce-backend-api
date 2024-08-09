package user

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public Route
	// ur := repo.NewUserRepo()
	// us := service.NewUserService(ur)
	// userHandlerNonDependency := controller.NewUserController(us)

	userController, _ := wire.InitUserRouterHandler()

	publicUserRouter := Router.Group("/user")
	{
		publicUserRouter.POST("/register", userController.Register)
		publicUserRouter.POST("/otp")
	}

	// Private Route
	privateUserRouter := Router.Group("/product")
	{
		privateUserRouter.GET("/getInfo")
	}
}
