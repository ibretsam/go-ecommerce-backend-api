package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public Route
	publicAdminRouter := Router.Group("/admin")
	{
		publicAdminRouter.POST("/login")
	}

	// Private Route
	privateAdminRouter := Router.Group("/admin/user")
	{
		privateAdminRouter.POST("/activateUser")
	}
}
