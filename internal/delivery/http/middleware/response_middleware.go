// delivery/http/middleware/response_middleware.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mot0x0/gopi/internal/delivery/http/response"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Written() {
			return
		}

		if resp, exists := c.Get(response.ResponseKey); exists {
			apiResponse := resp.(response.APIResponse)
			c.JSON(apiResponse.Status, apiResponse)
			return
		}

		statusCode := c.Writer.Status()
		defaultResponse := response.APIResponse{
			Success: statusCode < 400,
			Status:  statusCode,
			Message: http.StatusText(statusCode),
		}

		if len(c.Errors) > 0 {
			defaultResponse.Error = c.Errors.Last().Error()
		}

		c.JSON(statusCode, defaultResponse)
	}
}
