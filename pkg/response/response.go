package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // Status code
	Message string      `json:"message"` // Message
	Data    interface{} `json:"data"`    // Data
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	})
}
