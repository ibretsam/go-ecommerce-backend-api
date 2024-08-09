package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Private Route
	privateUserRouter := Router.Group("/product")
	{
		privateUserRouter.GET("/activateUser")
	}
}
