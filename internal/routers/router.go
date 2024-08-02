package routers

import (
	"fmt"
	c "go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> CC")
		c.Next()
		fmt.Println("After --> CC")
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middlewares.AuthMiddleware(), BB(), CC())

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}

	return r
}
