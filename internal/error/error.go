package error

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			errMeta := c.Errors.Last().Meta

			c.JSON(errMeta.(int), gin.H{
				"success": false,
				"error":   err.Error(),
			})
		}
	}
}
