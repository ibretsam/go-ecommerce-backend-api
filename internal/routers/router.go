package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "masterbi")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong" + name,
		"uid":     uid,
		"user":    []string{"masterbi1", "masterbi2"},
	})
}
