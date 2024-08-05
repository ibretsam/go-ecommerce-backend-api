package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public Route
	publicUserRouter := Router.Group("/user")
	{
		publicUserRouter.POST("/register")
		publicUserRouter.POST("/otp")
	}

	// Private Route
	privateUserRouter := Router.Group("/product")
	{
		privateUserRouter.GET("/getInfo")
	}
}
