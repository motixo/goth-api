// delivery/http/response/helper.go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const ResponseKey = "api_response"

type APIResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, message string, data interface{}) {
	c.Set(ResponseKey, APIResponse{
		Success: true,
		Status:  status,
		Message: message,
		Data:    data,
	})
	c.Status(status)
}

func Error(c *gin.Context, status int, message string, err error) {
	response := APIResponse{
		Success: false,
		Status:  status,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.Set(ResponseKey, response)
	c.Status(status)
}

func Created(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusCreated, message, data)
}

func OK(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusOK, message, data)
}

func BadRequest(c *gin.Context, message string, err error) {
	Error(c, http.StatusBadRequest, message, err)
}

func InternalError(c *gin.Context, message string, err error) {
	Error(c, http.StatusInternalServerError, message, err)
}

func Unauthorized(c *gin.Context, message string, err error) {
	Error(c, http.StatusUnauthorized, message, err)
}
